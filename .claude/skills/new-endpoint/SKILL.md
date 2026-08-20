---
name: new-endpoint
description: >
  Scaffold a new REST API endpoint in the ops-api service. Use when the user wants
  to add a new route, endpoint, or API operation including the handler, service,
  repository, and route wiring.
argument-hint: "[resource] [operation]"
---

# New API Endpoint

You are adding a new REST API endpoint to the ops-api (`cmd/api/`). This follows
the hexagonal architecture: handler (primary adapter) -> service (core) -> repository (secondary adapter).

The user will provide:
- **Resource name** (e.g., `session`, `role`) — use `$ARGUMENTS[0]` if provided
- **Operation** (e.g., `GetByAddress`, `List`) — use `$ARGUMENTS[1]` if provided

If any of these are missing, ask the user before proceeding.

## File Map

| Layer | File | Purpose |
|-------|------|---------|
| Route | `cmd/api/app/app.go` | Gin route registration |
| Handler | `cmd/api/adapters/handlers/<resource>_handler.go` | HTTP handler with Swagger annotations |
| Service interface | `cmd/api/core/ports.go` (Primary Ports) | Business operation interface |
| Service impl | `cmd/api/core/<resource>_service.go` | Business logic |
| Repo interface | `cmd/api/core/ports.go` (Secondary Ports) | Data access interface |
| Repo impl | `cmd/api/adapters/repositories/<resource>_repository.go` | GORM queries |
| Domain | `domain/<resource>.go` | Domain model (if new entity) |
| DI | `di/container.go` | Wire repo → service → handler, add handler as Container field |

Use the health handler as the simplest reference implementation.

## Step 1: Define the repository interface in `cmd/api/core/ports.go`

```go
// <Resource>Repository handles database queries for <resources>
type <Resource>Repository interface {
    FindBy<Field>(ctx context.Context, field string) (*domain.<Resource>, error)
}
```

Key patterns:
- All methods take `context.Context` as first argument
- Return `core.ErrRecordNotFound` when a record doesn't exist
- Use descriptive method names: `FindBy...`, `Create...`, `Update...`

## Step 2: Define the service interface in `cmd/api/core/ports.go`

```go
// <Resource>Service defines the business operations for <resources>
type <Resource>Service interface {
    Get<Resource>By<Field>(ctx context.Context, field string) (*domain.<Resource>, error)
}
```

## Step 3: Implement the repository

Create `cmd/api/adapters/repositories/<resource>_repository.go`:
- Unexported struct, exported constructor returning the interface
- Always use `r.db.WithContext(ctx)`
- Map `gorm.ErrRecordNotFound` to `core.ErrRecordNotFound`

## Step 4: Implement the service

Create `cmd/api/core/<resource>_service.go`:
- Unexported struct, exported constructor returning the interface
- Validate inputs, return domain errors (`NewValidationError`, `NewNotFoundError`, `NewInternalError`)
- Never return raw DB errors

## Step 5: Create the HTTP handler

Create `cmd/api/adapters/handlers/<resource>_handler.go`:
- Include Swagger annotations (`@Summary`, `@Param`, `@Success`, `@Failure`, `@Router`)
- Use `c.Request.Context()` when calling services
- Use `HandleError(c, h.log, err)` for error responses

## Step 6: Wire in `di/container.go`

Add the handler as a field in `Container`:
```go
type Container struct {
    // ... existing fields
    <Resource>Handler *handlers.<Resource>Handler
}
```

In `New()`, add:
```go
<resource>Repo := repositories.New<Resource>Repository(db)
<resource>Service := core.New<Resource>Service(<resource>Repo, log)
container.<Resource>Handler = handlers.New<Resource>Handler(<resource>Service, log)
```

## Step 7: Register the route in `cmd/api/app/app.go`

```go
router.GET("/<resources>/:<field>", container.<Resource>Handler.Get<Resource>By<Field>)
```

For auth-protected routes:
```go
router.GET("/<resources>/:<field>", authMiddleware, container.<Resource>Handler.Get<Resource>By<Field>)
```

For role-protected routes:
```go
router.GET("/<resources>/:<field>", authMiddleware, requireRole(RoleAdmin), container.<Resource>Handler.Get<Resource>By<Field>)
```

## Step 8: Write tests

Create `cmd/api/core/<resource>_service_test.go`:
- One function per scenario, not table-driven
- Use fake repositories from `cmd/api/testutil/`
- Use `testutil.StubLogger{}` for logging
- Test: happy path, not found, validation errors

## Step 9: Verify

```bash
make build
make lint
make swagger
make test
```

## Checklist

- [ ] Repository interface in `core/ports.go`
- [ ] Service interface in `core/ports.go`
- [ ] Repository implementation in `adapters/repositories/`
- [ ] Service implementation in `core/`
- [ ] HTTP handler in `adapters/handlers/` with Swagger annotations
- [ ] Handler field added to `di/container.go`
- [ ] Wiring in `di/container.go` `New()` function
- [ ] Route registered in `app/app.go`
- [ ] Service tests in `core/`
- [ ] Swagger docs regenerated
- [ ] Build + lint pass
