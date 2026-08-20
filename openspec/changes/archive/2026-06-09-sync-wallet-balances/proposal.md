## Why

Today the ops-api knows every user's wallet (`user_wallets`) and every on-chain token (`tokens`), but cannot answer the most basic question for product surfaces: "how much of token X does wallet Y hold?". Blockscout already indexes per-address balances in its own database (`address_current_token_balances`), and we already mirror the Blockscout `tokens` table into our schema via a cursor backfiller + `LISTEN/NOTIFY` trigger. We should reuse that exact integration pattern to sync wallet balances so the API can serve them (and live updates) without re-implementing on-chain scanning.

## What Changes

- Add a new `wallet_balances` table in ops-api that stores, per `(wallet_address, token_address)`, the current balance, last-known block, and `updated_at`. Addresses are normalized via `domain.NormalizeAddress`.
- Add a Blockscout-side trigger (`scripts/blockscout_balances_trigger.sql`) on `address_current_token_balances` that emits `pg_notify('balance_change', json)` on `INSERT/UPDATE`. Applied at startup like the existing token trigger, via `infrastructure/database/blockscout.go`.
- Add a `BlockscoutBalancesBackfiller` (cursor on `address_current_token_balances.id`) and a `BlockscoutBalancesListener` (LISTEN/NOTIFY with exponential-backoff reconnect), both filtering to addresses that belong to a known `user_wallets.rayls_address`. Both run inside the `worker` entrypoint.
- Publish a durable event `ops.wallet_balances.updated` to a new JetStream stream `WALLET_BALANCE_EVENTS`, and a live `ops.sse.wallet_balances` core-NATS event for the API fan-out — mirroring how tokens already work.
- Add HTTP endpoints under `/api/wallets/:address/balances` (list current balances for one wallet) and `/api/wallets/balances/stream` (SSE for live updates scoped to the caller's wallet). Auth-gated like the existing token stream.
- Wire the new repository, service, handler, indexers, stream, and SSE topic in `di/container.go`, `cmd/api/app/worker.go`, and `cmd/api/app/app.go`.

## Capabilities

### New Capabilities
- `wallet-balances`: per-wallet, per-token balance tracking sourced from Blockscout, exposed through REST and SSE.

### Modified Capabilities
<!-- None: existing capabilities (tokens, blockscout-indexer) are not changing their requirements; this change adds a parallel capability that reuses the same patterns. -->

## Impact

- **New code**: `cmd/api/indexer/blockscout_balances_backfill.go`, `cmd/api/indexer/blockscout_balances_listener.go`, `cmd/api/adapters/repositories/wallet_balance_repository.go`, `cmd/api/services/wallet_balance_service.go`, `cmd/api/adapters/handlers/wallet_balance_handler.go`, `cmd/api/adapters/handlers/wallet_balance_stream_handler.go`, `domain/wallet_balance.go`, `scripts/blockscout_balances_trigger.sql`.
- **Modified code**: `di/container.go` (wire repo/service/handler/indexers and add `WALLET_BALANCE_EVENTS` JetStream stream), `cmd/api/app/worker.go` (start balances backfiller + listener after the tokens ones), `cmd/api/app/app.go` (register routes), `infrastructure/database/blockscout.go` (apply the new trigger script), `cmd/api/messaging/manager.go` (no API change, just an additional stream registered via container).
- **New migration**: `migrations/00000X_wallet_balances.up.sql` / `.down.sql` for the `wallet_balances` table and `indexer_state` rows are reused (new cursor key, no schema change).
- **External dependency**: Requires `BLOCKSCOUT_DB_CONN` (already present) and assumes Blockscout schema has `address_current_token_balances(address_hash, token_contract_address_hash, value, block_number, value_fetched_at, id)`. The trigger is applied idempotently on every API startup, same pattern as the tokens trigger.
- **No breaking changes**: purely additive. Existing token endpoints, streams, and indexers are untouched.
