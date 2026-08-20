---
name: investigate
description: >
  Investigate bugs and issues in the ops-api service — authentication failures
  (SIWE or Web2), authorization errors, login flow problems, data issues, and
  infrastructure errors. Use when the user reports a 401/403 error, login failure,
  nonce issue, JWT problem, credential issue, role mismatch, "record not found"
  error, wrong API response, missing field, or any bug in the service. Also use
  when the user pastes an error log, asks "why is this failing", says "something
  looks wrong", or wants to trace how auth or data flows through the system.
argument-hint: '[endpoint or entity] [symptom]'
---

# Investigate Bug

You are investigating a bug in the ops-api. Your goal is to find the **root cause**,
not just identify which layer has the problem. Trace the request from entry point
through every layer until the point where the bug becomes visible.

## System Overview

Single-service REST API with SIWE authentication and role-based access control:

```
Client → Gin Router → Middleware (Auth/Role) → Handler → Service → Repository → DB
                                                    ↓
                                              Smart Contract (role resolution via RPC)
```

| Layer | Location | Role |
|-------|----------|------|
| **Routes** | `cmd/api/app/app.go` | Route registration, middleware assignment |
| **Middleware** | `cmd/api/middleware/` | Auth validation, query param validation |
| **Handlers** | `cmd/api/adapters/handlers/` | HTTP request/response, Swagger annotations |
| **Services** | `cmd/api/core/` | Business logic, validation, error mapping |
| **Repositories** | `cmd/api/adapters/repositories/` | GORM database queries |
| **DI Container** | `di/container.go` | All dependency wiring |
| **Domain** | `domain/` | Domain models |
| **Config** | `config/` | Viper-based configuration |

## Root Cause Investigation Method

Work backwards from the symptom to the source. At each layer, read the actual code —
don't guess based on file names alone.

### Step 1: Classify the issue

Before reading any code, identify the category:

| Category | Symptoms | Start investigating at |
|----------|----------|----------------------|
| **Auth (401)** | Login fails, JWT expired, cookie missing, nonce invalid, bad credentials | Middleware → Auth handler → Nonce/credential service |
| **Authorization (403)** | Role mismatch, permission denied, wrong role resolved | Role middleware → Role resolver → Smart contract / DB |
| **SIWE flow** | Signature verification fails, message format wrong, nonce expired | Auth service → SIWE verification logic |
| **Web2 auth flow** | Wrong password, account locked, credential validation fails | Auth service → Credential verification logic |
| **Data (404/wrong value)** | Record not found, wrong field value, missing data | Handler → Service → Repository → DB query |
| **Infrastructure (500)** | DB connection, RPC node timeout, config missing | DI container → Config → DB/RPC setup |

### Step 2: Trace backwards from the symptom

Start at the layer where the bug is visible and work toward the data source.

**Middleware layer** — is the request being rejected before reaching the handler?
- `cmd/api/middleware/` — auth middleware extracts JWT from cookie, validates claims
- Check: is the cookie being set? Is the JWT valid? Is the role claim present?

**Handler layer** — is the request parsed and dispatched correctly?
- `cmd/api/adapters/handlers/` — params extraction, response format
- Check: are URL params, query params, body parsed correctly?

**Service layer** — is the business logic correct?
- `cmd/api/core/` — validation, error types, domain logic
- Check: are inputs validated? Are errors mapped correctly?

**Repository layer** — is the database query correct?
- `cmd/api/adapters/repositories/` — GORM queries, JOINs
- Check: is the query selecting the right columns? Is `gorm.ErrRecordNotFound` mapped to `core.ErrRecordNotFound`?

**DI wiring** — are dependencies connected correctly?
- `di/container.go` — all wiring happens here
- Check: is the new handler/service/repo wired? Is the right interface satisfied?

**Smart contract layer** — is the on-chain call returning expected data?
- Role resolver calls the access control contract via RPC
- Check: is the contract address correct? Is the RPC endpoint reachable? Is the ABI up to date?

### Step 3: Identify the root cause

Once you find where the data diverges from what's expected, determine **why**:
- Is it a logic error in the code?
- Is it a configuration issue (wrong env var, missing config)?
- Is it an infrastructure issue (RPC node unreachable, DB down)?
- Is it a wiring issue (dependency not connected in DI container)?
- Is it a data model issue (wrong type, missing field, enum mismatch)?

Always explain the full causal chain: "Request hits X → middleware does Y → but Y fails because..."

## Auth Flow Investigation

The API supports two auth methods. Both produce the same JWT structure — downstream
middleware is auth-method-agnostic.

### SIWE Authentication (Web3)

```
1. GET  /auth/nonce?address=0x...  → NonceService.Generate()
                                   → Store nonce in DB with expiry
                                   → Return nonce to client

2. POST /auth/siwe/login           → AuthService.SiweLogin()
   Body: { message, signature }    → Parse SIWE message (EIP-4361)
                                   → Verify signature against address
                                   → Verify nonce matches stored nonce
                                   → Verify nonce not expired
                                   → RoleResolver.GetRole(address) via smart contract
                                   → Generate JWT with role claim
                                   → Set HTTP-only cookie ("Authorization")
```

### Standard Authentication (Web2)

```
1. POST /auth/login                → AuthService.Login()
   Body: { username, password }    → Verify credentials against DB
                                   → Resolve role from DB
                                   → Generate JWT with role claim
                                   → Set HTTP-only cookie ("Authorization")
```

### Shared (both flows)

```
3. POST /auth/logout               → Clear session/cookie

4. Protected routes                → RequireAuth middleware
                                   → Extract JWT from "Authorization" cookie
                                   → Validate JWT signature and expiry
                                   → Set user context (identifier, role)
                                   → RequireRole(role) checks role claim
```

**Common auth issues:**
| Symptom | Likely cause | Where to look |
|---------|-------------|---------------|
| 401 on SIWE login | Nonce expired or already used | Nonce service/repository, expiry config |
| 401 on Web2 login | Bad credentials or account state | Credential verification, user repository |
| 401 on protected route | JWT expired or cookie not sent | Middleware, cookie config (domain, secure, httpOnly) |
| 403 on protected route | Role doesn't match requirement | Role resolver, smart contract / DB role state |
| SIWE verification fails | Message format mismatch | SIWE parsing logic, EIP-4361 compliance |
| Login succeeds but role wrong | Contract/DB returns unexpected value | Role resolver, contract ABI, user record |

## Common Bug Patterns

| Symptom | Likely cause | Where to look |
|---------|-------------|---------------|
| "record not found" | Query condition wrong or data missing | Repository query, service error mapping |
| 401 Unauthorized | JWT/cookie issue | Auth middleware, cookie settings |
| 403 Forbidden | Role mismatch | Role middleware, role resolver |
| 500 Internal Server Error | Unhandled error, DB issue, RPC timeout | Service layer, DI wiring, config |
| Wrong field in response | DTO mapping or query issue | Handler response construction, repository SELECT |
| Missing field in response | Field not queried or not mapped | Repository query, domain model |
| Handler not reachable | Route not registered or middleware blocks | `app.go` routes, middleware chain |
| Dependency nil pointer | Not wired in DI container | `di/container.go` New() function |

## Key Files Reference

**Interfaces:** `cmd/api/core/ports.go` — all service and repository interfaces

**Error types:** `cmd/api/core/errors.go` — `NotFoundError`, `ValidationError`, `InternalError`

**Error handler:** `cmd/api/adapters/handlers/error_handler.go` — `HandleError()` maps domain errors to HTTP status

**DI container:** `di/container.go` — all dependency construction and wiring

**Routes:** `cmd/api/app/app.go` — route registration with middleware

**Domain models:** `domain/` — base model with UUID, timestamps

**Config:** `config/` — Viper with `ExperimentalBindStruct`

## Investigation Checklist

When investigating, read the actual code at each layer — don't just check file names.

- [ ] **Symptom clear**: entity, endpoint, expected vs actual behavior identified
- [ ] **Category identified**: auth (401), authorization (403), SIWE, data, infrastructure
- [ ] **Middleware layer**: auth middleware, role middleware, request rejection
- [ ] **Handler layer**: params extraction, response construction, Swagger match
- [ ] **Service layer**: business logic, validation, error types
- [ ] **Repository layer**: GORM query, JOIN conditions, error mapping
- [ ] **DI wiring**: dependency connected in container, interface satisfied
- [ ] **Smart contract**: RPC call, ABI, contract address, role returned
- [ ] **Config**: environment variables, Viper binding, required fields
- [ ] **Root cause**: full causal chain identified, not just the layer where it breaks
