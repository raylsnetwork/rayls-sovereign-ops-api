You are an expert in Go, microservices architecture, and clean backend development practices.

Go coding standards (naming, imports, testing patterns, git conventions, error messages) are in `.claude/rules/code-standards.md`.
This file covers **project-specific** structure, patterns, and tooling for ops-api.

## Project Structure

Single-service application with two entrypoints (`run` = API, `worker` = indexers):
- `cmd/api/` - main entrypoint
- `cmd/api/app/` - entrypoints (`run` = API server, `worker` = background indexers) and route registration
- `cmd/api/core/` - domain errors and port interfaces
- `cmd/api/adapters/handlers/` - HTTP handlers with Swagger annotations
- `cmd/api/adapters/repositories/` - repository implementations (GORM)
- `cmd/api/adapters/blockchain/` - on-chain services (AccessManager, RNContractFactory deploy, token mint/burn)
- `cmd/api/adapters/custody/` - custody (Rayls HSM) adapter
- `cmd/api/services/` - domain services (auth, provisioning, token permissions)
- `cmd/api/auth/` - SIWE / OAuth / JWT (go-pkgz/auth)
- `cmd/api/indexer/` - Blockscout (tokens) and AccessManager indexers (run in `worker`)
- `cmd/api/messaging/` - NATS/JetStream manager (durable publish/consume + core-NATS live for SSE)
- `cmd/api/sse/` - in-memory SSE hub
- `cmd/api/middleware/` - HTTP middleware (auth, query validation, encoding)
- `cmd/api/utils/` - shared HTTP utilities (error responses)
- `contracts/` - abigen **v2** contract bindings + `creator.go` (`EnsureCode`)
- `di/` - dependency injection container
- `domain/` - domain models
- `config/` - configuration management (Viper)
- `logger/` - structured logging (slog)
- `infrastructure/database/` - DB connection and migrations
- `migrations/` - SQL migration files for a chain's own database (golang-migrate)
- `migrations-identity/` - SQL migrations for the shared identity database (users, providers,
  login wallets, `user_signup_details`); applied by `NewIdentity`, never by the per-chain container
- `withstack/` - error wrapping with stack traces (cockroachdb/errors)
- `flags/` - CLI flag definitions (Cobra)

## DI Container (`di/container.go`)

All dependency wiring happens in the container, **not** in `app.go`:
- `New(configPath)` builds: infrastructure → repositories → services → handlers
- `app.go` only pulls handlers from the container and registers routes
- Add new repos/services/handlers as fields in `Container` struct

## Gin Framework

- Handlers in `adapters/handlers/` with Swagger annotations (`@Summary`, `@Param`, `@Success`, `@Failure`)
- Middleware for authentication and query validation
- `HandleError()` maps domain errors to HTTP status codes
- Group routes by resource

## GORM and Database

- Custom `Model` struct with UUID primary key and UTC timestamp hooks (`BeforeCreate`/`BeforeUpdate`)
- Migrations in `/migrations` named: `{timestamp}_{description}.{up|down}.sql`

## Error Handling

- `cockroachdb/errors` for stack traces via `withstack.Wrap()`
- Propagate errors with context, never swallow silently
- `HandleError()` maps domain errors to HTTP status codes in handlers

## Smart Contracts / On-chain

Contract bindings in `contracts/` are **abigen v2** (`Pack*`/`Unpack*`/`Instance`; no v1 address-bound
methods). On-chain services live in `cmd/api/adapters/blockchain/` (`RaylsAccessManagerService`,
`RaylsContractFactoryService` for deploy, `RaylsTokenService` for mint/burn).

- **Read call:** `data := b.PackX(args)` → `client.CallContract(ctx, ethereum.CallMsg{To:&addr, Data:data}, nil)` → `b.UnpackX(out)`
- **Tx:** `b.PackX(args)` → `types.NewTransaction(...)` → `MarshalBinary` → `custody.SignAndTransact` → poll receipt
- **Events:** `b.UnpackXEvent(&log)`
- `contracts.EnsureCode(ctx, client, addr)` validates bytecode exists (v2 has no address-bound constructor)
- Resolve contracts via `DeploymentProxyRegistry` by string key (`"RaylsAccessManager"`, `"RNContractFactory"`); address in `DEPLOYMENT_PROXY_REGISTRY_ADDR`
- Services reuse one `*ethclient.Client`; **signing is always delegated to custody (HSM)** — the key never leaves the HSM

## Addresses

Canonicalize EVM addresses with `domain.NormalizeAddress` (lowercase + `0x`) before storing/querying.
The token repository and Access Manager lookups rely on this to avoid duplicate rows / case mismatches.

## Tokens

Endpoints under `/api/tokens`: `POST` (deploy), `GET` (list), `GET /:address` (details),
`GET /:address/permissions`, `POST /:address/mint|burn`, `GET /api/tokens/stream` (SSE). Frontend contract
is documented in `docs/frontend-tokens-guide.md`.

## Wallet Balances

Endpoints under `/api/wallets`: `GET /:address/balances` (per-token balances for a wallet),
`GET /balances/stream` (SSE of live updates scoped to the caller's wallet). Frontend contract is in
`docs/frontend-balances-guide.md`.

Balances are synced from Blockscout's `address_current_token_balances` table by a cursor-based backfiller
+ `LISTEN/NOTIFY` listener (`cmd/api/indexer/blockscout_balances_*.go`) — same pattern as the tokens
indexer. The Blockscout-side trigger is installed automatically on API startup via
`infrastructure/database/blockscout.go:ApplyBlockscoutBalancesTrigger`, sourced from
`scripts/blockscout_balances_trigger.sql`. Durable events flow through the `WALLET_BALANCE_EVENTS`
JetStream stream (subject `ops.wallet_balances.>`); live SSE fan-out uses core NATS on
`ops.sse.wallet_balances`. Only addresses present in `user_wallets` are stored or broadcast.

- `domain.TokenStatus` includes the off-chain `TokenStatusInternal` (10): API-deployed tokens are stored
  Internal and **stay** internal (not promoted). `Label()` → string; responses expose `statusLabel` and
  `ercStandardLabel` (`RAYLS_ERC20`…). Deploy stores `issuerId` = chainId (the PN) and passes `resourceId` = `bytes32(0)`.
- Mint/burn build calldata **per standard** (no token bindings) and require the AM permission.
- Token sync: the Blockscout indexer (in `worker`) upserts tokens + publishes events; SSE relays a curated
  live event (worker publishes over **core NATS**, API fans out via the `cmd/api/sse` hub; JetStream is used
  for durable events).
- **Access Manager:** `am_*` tables are populated by the AccessManager indexer. `TokenPermissionService`
  derives a wallet's callable functions (global + contract-scoped roles ∩ `am_function_permissions`);
  selector→name map in `cmd/api/services/token_functions.go`.

## Configuration

- Environment variables via Viper's `ExperimentalBindStruct`
- Fallback: `.env` files
- `mapstructure` struct tags for field mapping
- `validator` library for required fields
- On-chain/infra vars: `BLOCKCHAIN_RPC_URL`, `DEPLOYMENT_PROXY_REGISTRY_ADDR`, `NATS_URL`, `BLOCKSCOUT_DB_CONN`, `INSTANCE_NAME` (gates on-chain services / indexers / SSE when set)

## Logging

- `slog` with colored text handler (dev) / JSON handler (production)
- Log levels: Debug, Info, Warn, Error

## Naming Conventions

- Receiver names: short lowercase (`r`, `p`, `t`, `s`)
- Methods: descriptive verbs (`FindByUsername`, `GetByResourceId`)

## Build Commands

```bash
make build      # Build the API service
make run        # Build and run the API service
make lint       # Run golangci-lint
make swagger    # Regenerate OpenAPI docs
make test       # Run all tests
```

## Running the API

The binary uses Cobra subcommands (separate processes, both build the same `di.Container`):

```bash
./build/rayls-ops-api run    --config config/.env   # HTTP API (port 8080) + SSE stream
./build/rayls-ops-api worker --config config/.env   # background indexers (Blockscout tokens, AccessManager)
```

- `--config` - path to `.env` file (optional; without it, Viper reads from OS environment variables)

## Custody Service

The custody service (Rayls HSM) is a .NET 8 ASP.NET Core API living at
`$HOME/project-dir/rayls-privacy-custody-light/`, a sibling of this project directory (mounted into the
`custody` container at `/app`).

When you need to understand the custody API contract (routes, request/response shapes), read the source directly:
- **Routes**: `src/Rayls.Custody.HSM.DTO/Const/CustodyHSMRoutes.cs`
- **Controllers**: `src/Rayls.Custody.HSM.API/Controllers/`
- **DTOs**: `src/Rayls.Custody.HSM.DTO/` (snake_case when `[JsonObject(NamingStrategyType = typeof(SnakeCaseNamingStrategy))]` is present, PascalCase otherwise)

Key contracts:
- **Auth**: `POST /api/auth` `{"ApiKey":"..."}` → `{"Token":"..."}` — exchange API key for Bearer JWT
- **Create wallet**: `POST /api/wallet` `{"type":"KEYSTORE_V3","address_quantity":1}` → `[{"id":"...","address":"..."}]`
- **Sign/send tx**: `POST /api/transaction` `{"WalletId":"...","Transaction":{"Data":"0x...","Value":"0x0"}}` → `{"TxHash":"..."}`
- **Wallet by address**: `GET /api/wallet/address/{address}` → `{"id":"..."}`
- **Health**: `GET /api/health` → `{"status":"healthy"}`

Our adapter is at `cmd/api/adapters/custody/rayls_hsm.go`. It caches the Bearer token and re-authenticates on 401.
