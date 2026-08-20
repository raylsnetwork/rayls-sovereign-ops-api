## Context

The ops-api already ingests token contract metadata from Blockscout using a two-arm pattern: a cursor-based **backfiller** that catches up on rows missed while we were offline, and a **listener** that uses `LISTEN/NOTIFY` to react to changes in near real time. Both are wired into the `worker` entrypoint, publish events to a JetStream stream, and also fan out a curated live event to the API via core NATS so the SSE hub can broadcast to connected clients. The Blockscout-side trigger is owned by us — we install it on every API startup via `infrastructure/database/blockscout.go`.

Wallet balances are a natural fit for the same pattern. Blockscout's stock schema includes `address_current_token_balances`, which is updated as the indexer observes transfers. We already know every "interesting" address because we created the wallet ourselves and stored it in `user_wallets.rayls_address`. The product needs both an at-rest read API (per-wallet balances) and a live stream for UI surfaces that show holdings updating.

Key constraints already encoded in the codebase:
- `BLOCKSCOUT_DB_CONN` is required for any Blockscout reads; if absent, indexers are gated off.
- `pg_notify` payloads must stay under 8 KB; `value` (uint256) can be very large, so we cast it to text and accept the size.
- All EVM addresses are normalized via `domain.NormalizeAddress` before storage and lookup.
- Signing/custody is unrelated here — we only read; no HSM calls.

## Goals / Non-Goals

**Goals:**
- Maintain a `wallet_balances` table keyed on `(wallet_address, token_address)` that is eventually consistent with Blockscout for every address in `user_wallets`.
- Serve current balances via `GET /api/wallets/:address/balances` and live updates via `GET /api/wallets/balances/stream`.
- Reuse the existing backfiller + LISTEN/NOTIFY + JetStream + SSE pattern with no architectural divergence from the tokens pipeline.
- Be safe to restart: cursored backfill, deduplicated publishes, idempotent trigger installation, idempotent upserts.

**Non-Goals:**
- Historical balance time-series. We store only the current balance plus the source block number.
- Native (ETH/RBC) balance — only ERC-20/721/1155 balances exposed by Blockscout's token-balances tables.
- Cross-chain balances. Single chain, single Blockscout instance, gated by `INSTANCE_NAME` like the tokens pipeline.
- Pushing balances to clients that do not own the wallet — the SSE filter is scoped to the authenticated user.

## Decisions

### Decision 1: Reuse the tokens "backfill + LISTEN/NOTIFY" pattern verbatim

We add two new files mirroring `blockscout_backfill.go` and `blockscout_listener.go`. Same cursor mechanism (stored in `indexer_state`), same reconnect/backoff, same publisher with deterministic MsgID, same dual publish (JetStream durable + core-NATS live).

**Alternatives considered:**
- *Polling JSON-RPC `balanceOf` for every (wallet, token) pair*: O(N×M) RPC calls, no good live signal, duplicates work Blockscout already does. Rejected.
- *Subscribing directly to Blockscout's Kafka/RabbitMQ output*: Blockscout's transport is optional and not deployed in our infra. The `address_current_token_balances` table is universally present. Rejected as premature coupling.
- *Logical replication slot on Blockscout*: heavier ops cost (replication user, slot lifecycle, restart semantics), no proven advantage over `pg_notify` for our event volume. Rejected.

### Decision 2: Filter to known wallets at ingestion, not at read time

Both the backfiller and the listener look up `address_hash` against `user_wallets` and only upsert if a match exists. This keeps `wallet_balances` lean and reads cheap.

**Alternatives considered:**
- *Storing every balance Blockscout knows*: wastes storage and bandwidth — Blockscout indexes the whole chain, the vast majority of addresses are not our users.
- *Read-time join only*: avoids the filter, but then the listener publishes events for every address change in the chain into our JetStream stream, which is a denial-of-service vector and a noisy SSE source.

The trade-off: when a brand-new wallet is created in `user_wallets`, its existing on-chain balances are not visible until either the next backfill sweep picks them up or the next on-chain transfer fires the trigger. Documented in **Risks** below.

### Decision 3: Trigger fires on `INSERT` and on `UPDATE OF value`

We deliberately limit the `UPDATE` trigger to the `value` column to avoid spurious notifications when Blockscout rewrites cosmetic columns (`value_fetched_at`) without a real balance change.

**Alternatives considered:**
- *Trigger on full row update*: noisier, larger NATS volume, no extra information.
- *Trigger on `value` OR `block_number`*: `block_number` changes whenever `value` changes; redundant.

### Decision 4: Backfill cursor on `id`, not on `(updated_at, address_hash)`

`address_current_token_balances` has a monotonically increasing `id` (bigserial). A single-column cursor is simpler and indexes naturally. The tokens backfill uses `(inserted_at, contract_address_hash)` because Blockscout's `tokens` table doesn't have a serial id we can rely on; balances do.

### Decision 5: New JetStream stream `WALLET_BALANCE_EVENTS`

Separate from `TOKEN_EVENTS` so retention, consumer groups, and ops dashboards can evolve independently. Same retention policy (`WorkQueue`, 1-hour dedup) for symmetry.

### Decision 6: SSE filtering happens in the handler, not in the publisher

The listener publishes one `ops.sse.wallet_balances` event per balance change. The stream handler subscribes once and filters per connection by comparing event `walletAddress` against the caller's `user_wallets` set. Keeps the publish-side stateless; matches how the token SSE already works.

### Decision 7: Domain model and labels

`domain.WalletBalance` carries `WalletAddress`, `TokenAddress`, `Balance string` (decimal, full precision — never floats), `BlockNumber uint64`, plus `Model` for UUID PK and timestamps. The handler joins with `tokens` to enrich the response with `tokenSymbol`, `tokenName`, `decimals` so the frontend doesn't need a second round-trip.

## Risks / Trade-offs

- **New wallet has no balances until next sweep** → Mitigation: when `POST /api/wallets` (or whatever creates a `user_wallets` row) succeeds, enqueue a one-shot backfill scan filtered to that single address (`address_current_token_balances WHERE address_hash = ?`). Cheap, bounded.
- **`pg_notify` payload > 8 KB if `value` is very large** → Mitigation: `value` for ERC20 fits comfortably under 8 KB as text. For ERC1155 with extreme token IDs this could approach the limit; the listener falls back to a `SELECT` by `id` if the payload is missing fields, matching the tokens listener's defensive deserialization.
- **Listener and backfiller race on the same row** → Mitigation: upsert is keyed on `(wallet_address, token_address)` with `block_number` used as a tiebreaker — we only overwrite when the incoming `block_number` is `>=` the stored one, so a stale backfill cannot regress a fresher listener update.
- **JetStream dedup window of 1 hour is too short for slow re-syncs** → Acceptable: the downstream consumer is idempotent (upsert into the same table), and a duplicate update is a no-op.
- **Wallet renames / address rotation** → Out of scope. `user_wallets.rayls_address` is currently immutable; if that changes, both pipelines (tokens and balances) need a coordinated re-architecture.

## Migration Plan

1. Ship the migration adding `wallet_balances` and an index on `wallet_address`. Reversible via the `.down.sql`.
2. Deploy the new code with the listener and backfiller behind a feature gate keyed off `INSTANCE_NAME` being set (same gate the tokens indexer uses).
3. On first boot the trigger is installed against Blockscout; the backfiller then walks the entire `address_current_token_balances` table once, populating only rows whose `address_hash` is in `user_wallets`. Expected one-time cost: bounded by the size of Blockscout's table.
4. Once steady-state is observed (cursor at HEAD, listener idle), expose the endpoints publicly in the frontend.

**Rollback**: disable the worker entrypoint or drop `INSTANCE_NAME`; the trigger is harmless if no one is listening, and can be dropped with the included `.down.sql` for the trigger if needed. The `wallet_balances` table can be left in place — no consumer outside ops-api reads it.

## Open Questions

- *Do we need per-token granular SSE topics (e.g., one subject per token) or is one stream filtered per wallet sufficient?* — Going with one stream for now; revisit if a single wallet ever holds thousands of tokens with high-frequency updates.
- *Should the read endpoint paginate?* — Skipped for v1: typical wallets hold tens of tokens. Add pagination only if real data shows we need it.
