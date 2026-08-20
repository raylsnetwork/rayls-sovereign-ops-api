# wallet-balances Specification

## Purpose

Provide per-wallet, per-token balance tracking sourced from Blockscout, exposed to clients through a REST API and a live SSE stream. Balances are synced for every address present in `user_wallets` and surfaced both as a list (all balances for a wallet) and as a single-resource lookup (balance for a specific wallet/token pair).

## Requirements

### Requirement: Wallet balance ingestion from Blockscout

The system SHALL maintain a local `wallet_balances` table that mirrors per-address per-token balances for every address present in `user_wallets`, sourced from Blockscout's `address_current_token_balances` table. Both the wallet address and the token contract address SHALL be normalized via `domain.NormalizeAddress` (lowercase, `0x`-prefixed) before storage and lookup. A balance row SHALL be uniquely identified by the composite key `(wallet_address, token_address)`.

#### Scenario: Backfill discovers a new balance

- **WHEN** the `BlockscoutBalancesBackfiller` runs and finds a row in Blockscout's `address_current_token_balances` whose `address_hash` matches an existing `user_wallets.rayls_address`
- **THEN** the system upserts a corresponding row into `wallet_balances` with the normalized wallet address, normalized token address, the on-chain balance, and the source `block_number`, and advances the backfill cursor past that Blockscout row.

#### Scenario: Backfill ignores balances for unknown wallets

- **WHEN** a row in Blockscout's `address_current_token_balances` has an `address_hash` that does not exist in `user_wallets`
- **THEN** the system SHALL skip the row, MUST NOT insert into `wallet_balances`, and MUST still advance the backfill cursor past it.

#### Scenario: Live update via LISTEN/NOTIFY

- **WHEN** Blockscout fires the `balance_change` notification for an address that exists in `user_wallets`
- **THEN** the `BlockscoutBalancesListener` SHALL upsert the new balance, publish `ops.wallet_balances.updated` to the `WALLET_BALANCE_EVENTS` JetStream stream, and publish a curated live event on the core-NATS subject `ops.sse.wallet_balances`.

#### Scenario: Listener reconnects on Blockscout disconnect

- **WHEN** the LISTEN/NOTIFY connection to Blockscout drops
- **THEN** the listener SHALL reconnect with exponential backoff starting at 1 second and capped at 30 seconds, matching the existing token listener behaviour.

### Requirement: Wallet balance HTTP API

The system SHALL expose authenticated HTTP endpoints to read wallet balances and to stream live updates. Wallet addresses in path parameters SHALL be normalized before lookup. The endpoints SHALL return only balances for wallets that exist in `user_wallets`.

#### Scenario: List balances for a known wallet

- **WHEN** an authenticated caller issues `GET /api/wallets/{address}/balances` for an address that exists in `user_wallets`
- **THEN** the system SHALL return `200 OK` with a JSON array of balance entries, each containing `walletAddress`, `tokenAddress`, `tokenSymbol`, `tokenName`, `decimals`, `balance` (decimal string), `blockNumber`, and `updatedAt`.

#### Scenario: List balances for an unknown wallet

- **WHEN** an authenticated caller issues `GET /api/wallets/{address}/balances` for an address that is not in `user_wallets`
- **THEN** the system SHALL return `404 Not Found` with a domain error indicating the wallet is unknown.

#### Scenario: Stream live balance updates

- **WHEN** an authenticated caller opens `GET /api/wallets/balances/stream`
- **THEN** the system SHALL return a `text/event-stream` connection, send a 25-second heartbeat, and SHALL forward every `ops.sse.wallet_balances` event whose `walletAddress` belongs to the caller as an SSE `message` event with the same JSON payload published by the listener.

### Requirement: Single wallet-token balance lookup

The system SHALL expose `GET /api/wallets/{address}/balances/{tokenAddress}` returning the balance for a specific (wallet, token) pair. Both path parameters SHALL be normalized via `domain.NormalizeAddress` before lookup. The response body SHALL be a single balance object with the same shape as one entry of the list endpoint, enriched with `tokenSymbol`, `tokenName`, and `decimals`.

#### Scenario: Known pair returns the balance

- **WHEN** an authenticated caller issues `GET /api/wallets/{address}/balances/{tokenAddress}` for a wallet present in `user_wallets` that holds a balance for the given token
- **THEN** the system SHALL return `200 OK` with a JSON object containing `walletAddress`, `tokenAddress`, `tokenSymbol`, `tokenName`, `decimals`, `balance` (decimal string), `blockNumber`, and `updatedAt`.

#### Scenario: Unknown wallet returns 404

- **WHEN** an authenticated caller issues `GET /api/wallets/{address}/balances/{tokenAddress}` for an address that is not in `user_wallets`
- **THEN** the system SHALL return `404 Not Found` with a domain error indicating the wallet is unknown, and MUST NOT query the balances table.

#### Scenario: Known wallet but no balance for the token returns 404

- **WHEN** an authenticated caller issues `GET /api/wallets/{address}/balances/{tokenAddress}` for a wallet present in `user_wallets` that holds no row for the given token
- **THEN** the system SHALL return `404 Not Found` with a domain error identifying the missing `wallet balance` resource by `wallet/token`.

### Requirement: Blockscout trigger installation

The system SHALL install an idempotent PostgreSQL trigger on Blockscout's `address_current_token_balances` table that emits `pg_notify('balance_change', json)` on `INSERT` and on `UPDATE` of the `value` column. The trigger SHALL be applied at API startup using the same mechanism as the existing `token_change` trigger.

#### Scenario: Trigger applied at startup

- **WHEN** the `run` entrypoint boots with a valid `BLOCKSCOUT_DB_CONN`
- **THEN** the system SHALL execute the trigger script against Blockscout, succeed when the trigger already exists, and log a single info entry recording installation or no-op.

#### Scenario: Trigger payload shape

- **WHEN** a Blockscout `address_current_token_balances` row is inserted or its `value` updated
- **THEN** the emitted notification payload SHALL be a JSON object containing `op` (`"INSERT"` or `"UPDATE"`), `address_hash`, `token_contract_address_hash`, `value` (text), and `block_number`.

### Requirement: Wallet balance event publishing

The system SHALL register a JetStream stream named `WALLET_BALANCE_EVENTS` filtering on subject `ops.wallet_balances.>`, configured for work-queue retention and a 1-hour deduplication window, mirroring the existing `TOKEN_EVENTS` stream. Each published event SHALL use a deterministic MsgID derived from the subject and payload hash so that re-processing the same Blockscout row does not duplicate the event within the dedup window.

#### Scenario: Re-publishing the same balance is deduplicated

- **WHEN** the backfiller restarts and re-processes a Blockscout row whose balance has not changed
- **THEN** the second publish to `ops.wallet_balances.updated` SHALL be deduplicated by JetStream within the 1-hour window and SHALL NOT trigger a duplicate downstream delivery.
