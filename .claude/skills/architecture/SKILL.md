---
name: architecture
description: >
  Project architecture knowledge for the ops-api. Use when the user asks
  about project structure, where to put new code, how auth works, how the
  DI container is organized, or when making architectural decisions.
---

# OPS API Architecture

## Overview

The ops-api is a single-service REST API for the Rayls OPS dashboard. It supports
two authentication methods:
- **SIWE (Sign-In with Ethereum)** — Web3 wallet-based auth with on-chain role resolution
- **Standard Web2 auth** — username/password with traditional credential management

Both flows issue JWTs stored in HTTP-only cookies and share the same role-based
middleware for authorization.

## Single Service

| Service | Entry | Port | Responsibility |
|---------|-------|------|----------------|
| **API** | `cmd/api/main.go` | 8080 | REST API (Gin) with SIWE + Web2 auth, role-based middleware, health checks |

## Hexagonal Architecture

```
cmd/api/
  main.go              # Entrypoint (Cobra CLI)
  app/
    app.go             # Route registration only (pulls handlers from DI container)
    cmd.go             # Cobra command definition
  core/
    ports.go           # All interfaces (primary + secondary ports)
    errors.go          # Domain errors (NotFoundError, ValidationError, InternalError)
    *_service.go       # Business logic implementations
  adapters/
    handlers/          # HTTP handlers with Swagger annotations
    repositories/      # GORM-based PostgreSQL implementations
  middleware/          # HTTP middleware (auth, query validation)
  utils/              # Shared HTTP utilities (error responses)
  testutil/           # Fakes, stubs for testing
  docs/               # Generated Swagger docs

# Root-level shared packages
di/                   # DI container — ALL wiring happens here
domain/               # Domain models (UUID base model)
config/               # Viper config with ExperimentalBindStruct
logger/               # slog structured logging
infrastructure/
  database/           # DB connection + migrations
migrations/           # SQL migration files (golang-migrate)
withstack/            # Error wrapping with stack traces
flags/                # CLI flag definitions
```

## DI Container (`di/container.go`)

Unlike governance-api's `infrastructure.SetupInfrastructure()` pattern, ops-api uses
a **full DI container**:

```go
type Container struct {
    Config        *config.Config
    DB            *gorm.DB
    Logger        logger.Logger
    HealthHandler *handlers.HealthHandler
    // Add new handlers here as the service grows
}
```

- `New(configPath)` builds everything: config → DB → migrations → repos → services → handlers
- `app.go` only calls `container.HealthHandler.HealthCheck` etc. to register routes
- **All new dependencies go in Container** — never wire directly in `app.go`

## Dependency Flow

```
main.go → cmd.go → app.go
  → di.New(configPath)        # Builds entire dependency graph
    → config.LoadConfig()
    → database.Connect()
    → database.RunMigrations()
    → repositories.New*()
    → core.New*Service()
    → handlers.New*Handler()
  → app.go registers routes using container.XxxHandler
```

All dependencies point **inward**: adapters depend on core interfaces, never the reverse.

## Authentication Flows

### SIWE (Web3)

```
1. GET  /auth/nonce?address=0x...  → Generate nonce, store in DB with expiry
2. POST /auth/siwe/login           → Verify SIWE message + signature (EIP-4361)
                                   → Resolve role from smart contract
                                   → Issue JWT with role claim in HTTP-only cookie
3. POST /auth/logout               → Invalidate session
```

### Standard Web2

```
1. POST /auth/login                → Verify username/password credentials
                                   → Resolve role from DB
                                   → Issue JWT with role claim in HTTP-only cookie
2. POST /auth/logout               → Invalidate session
```

### Shared Middleware (both flows)

```
4. Middleware: RequireAuth         → Extract JWT from cookie, validate, set user context
5. Middleware: RequireRole(role)   → Check role claim against required role
```

Both flows produce the same JWT structure. Downstream middleware and handlers are
auth-method-agnostic — they only inspect the JWT claims.

## Smart Contract Integration

- Access control contract deployed on private network hub
- Go bindings generated via `abigen`
- `RoleResolver` interface abstracts contract calls:
  ```go
  type RoleResolver interface {
      GetRole(ctx context.Context, address string) (Role, error)
  }
  ```
- Error handling with retry and timeout for RPC calls

## Key Patterns

- **Interface-driven**: All cross-boundary calls go through interfaces in `core/ports.go`
- **DI container**: All wiring in `di/container.go`, not scattered across `app.go`
- **Receiver names**: Short lowercase (`r`, `s`, `h`)
- **Error handling**: `cockroachdb/errors` + `withstack.Wrap()` for stack traces
- **Domain errors**: `NewValidationError`, `NewNotFoundError`, `NewInternalError`
- **Config**: Viper `ExperimentalBindStruct()` with `mapstructure` tags

## Build and Test

```bash
make build      # Build the API binary
make run        # Build and run
make lint       # golangci-lint
make swagger    # Regenerate OpenAPI docs
make test       # Run all tests
```
