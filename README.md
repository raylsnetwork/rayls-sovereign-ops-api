<div align="center">

# Rayls Privacy Ops API

**Operations service (API + worker) for a Rayls privacy chain — token management, wallet balances, access-manager indexing, and custody-backed signing.**

[![License: Apache 2.0][license-badge]][license-url]
[![Go][go-badge]][go-url]

[![Discord][discord-badge]][discord-url]
[![X][x-badge]][x-url]
[![LinkedIn][linkedin-badge]][linkedin-url]
[![YouTube][youtube-badge]][youtube-url]

[Quick start](#quick-start) | [Running](#running-the-api) | [Auth](#authentication) | [API docs](#api-docs)

</div>

## What is this?

The Rayls Ops API is the operations service for a Rayls privacy chain. It exposes an HTTP API (Gin) and a background `worker`, both built from the same DI container:

- **Tokens** — deploy, list, mint/burn, and stream (SSE) tokens on the chain.
- **Wallet balances** — per-wallet balances synced from Blockscout, with live SSE updates.
- **Access Manager** — indexes on-chain roles/permissions and derives each wallet's callable functions.
- **Auth** — Google / Microsoft OAuth and SIWE (Sign-In with Ethereum), issuing JWTs.
- **Custody** — all signing is delegated to the Rayls HSM (keys never leave custody).

On-chain signing goes through the custody (HSM) service; the API resolves contracts via a `DeploymentProxyRegistry`.

## Prerequisites

- Go 1.24.2+
- PostgreSQL
- Docker (for the containerized dev stack)

## Quick Start

```bash
cp config/.env.example config/.env
# fill in database connection, OAuth credentials, JWT secret, etc.
```

See `config/.env.example` for all available settings.

### Development (hot reload + debugger)

Uses `Dockerfile.dev` with [Air](https://github.com/air-verse/air) for hot reload and [Delve](https://github.com/go-delve/delve) for remote debugging:

```bash
docker compose -f docker-compose.dev.yml up --build
```

This starts **ops-api** on `http://localhost:8080`, **PostgreSQL** on `localhost:5432`, and the **Delve debugger** on `localhost:2345`. Set `GOOGLE_CLIENT_ID` / `GOOGLE_CLIENT_SECRET` / `MICROSOFT_CLIENT_ID` / `MICROSOFT_CLIENT_SECRET` in `docker-compose.dev.yml` (or as environment variables) to enable OAuth.

For the full local stack against a RayUp-provisioned chain (ops-api + worker + custody + playground), see [`docs/flows.md`](docs/flows.md) and run `./start_dev.sh` (needs a local RayUp control plane on k3d).

### Development with Blockscout (shared postgres)

To reuse Blockscout's shared postgres instead of a separate one, the `docker-compose.blockscout.yml` override swaps the local postgres for Blockscout's `shared-db`. Bring up the shared stack first, then:

```bash
docker compose -f docker-compose.blockscout.yml up --build
```

### Production

Build and run the minimal (`scratch`-based) production image:

```bash
docker build -t rayls-ops-api .
docker run -p 8080:8080 \
  -e DATABASE_CONNECTIONSTRING="host=<host> port=5432 user=<user> password=<pass> dbname=<db> sslmode=disable" \
  -e JWT_SECRET="<secret>" \
  -e BASE_URL="https://your-domain.com" \
  -e CORSURLS="https://your-frontend.com" \
  rayls-ops-api
```

The entrypoint is `/app/rayls-ops-api run`; pass `--config` to use a mounted `.env` file.

## Make Commands

```bash
make build      # Build the API binary
make run        # Build and run the API
make lint       # Run golangci-lint
make swagger    # Regenerate OpenAPI docs
make test       # Run all tests
```

## Running the API

The binary uses Cobra subcommands (two processes, one DI container):

```bash
./build/rayls-ops-api run    --config config/.env   # HTTP API (port 8080) + SSE
./build/rayls-ops-api worker --config config/.env   # background indexers (tokens, AccessManager)
```

## Bootstrap

After first start, create the initial admin user:

```bash
curl -X POST http://localhost:8080/admin/bootstrap \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com"}'
```

On success the API returns `201 Created` with the provisioned HSM wallet address.

> **Note:** This endpoint can be called only once — if any user exists it returns `409 Conflict` (`BOOTSTRAP_ALREADY_COMPLETED`). It requires no authentication and should be protected at the network/infrastructure level in production.

## API Docs

Swagger UI is at `http://localhost:8080/swagger/index.html` while the server is running.

## Authentication

Three methods are supported:

- **Google OAuth** — `/auth/google/login`
- **Microsoft OAuth** — `/auth/microsoft/login`
- **SIWE (Sign-In with Ethereum)** — `/auth/siwe/login`

See [`docs/architecture.md`](docs/architecture.md) and [`docs/frontend-auth-guide.md`](docs/frontend-auth-guide.md) for details.

### Google OAuth Setup

1. Go to the [Google Cloud Console](https://console.cloud.google.com/) and create/select a project.
2. Under **APIs & Services > OAuth consent screen**, choose the user type, fill in the app details, and add the `userinfo.profile` and `userinfo.email` scopes (plus test users while in "Testing").
3. Under **APIs & Services > Credentials**, create an **OAuth client ID** (Web application) and add `http://localhost:8080/auth/google/callback` to the Authorized redirect URIs (use your production callback in production).
4. Copy the Client ID and Client Secret into your `.env`:
   ```
   GOOGLE_CLIENT_ID=your-client-id.apps.googleusercontent.com
   GOOGLE_CLIENT_SECRET=your-client-secret
   ```

## Documentation

- [Architecture](docs/architecture.md) — services, DI, on-chain integration
- [Flows](docs/flows.md) — local dev flow (RayUp) and provisioning
- [Frontend: Auth](docs/frontend-auth-guide.md) · [Tokens](docs/frontend-tokens-guide.md) · [Balances](docs/frontend-balances-guide.md)

## Contributing

We are not accepting external contributions at this time — see [CONTRIBUTING.md](./CONTRIBUTING.md). Please also read our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Security

To report a security vulnerability, see [SECURITY.md](./SECURITY.md) — please do not open a public issue.

## License

Licensed under the Apache License, Version 2.0 — see [LICENSE](./LICENSE).

This project links third-party libraries that remain under their own licenses; notably [go-ethereum](https://github.com/ethereum/go-ethereum) under the LGPL-3.0 (library packages only) and the HashiCorp `errwrap` / `go-multierror` libraries under the MPL-2.0. See [NOTICE](./NOTICE).

Copyright 2026 Rayls Core Ltd.

[license-badge]: https://img.shields.io/badge/License-Apache_2.0-blue.svg
[license-url]: ./LICENSE
[go-badge]: https://img.shields.io/badge/Go-1.24.2-00ADD8?logo=go&logoColor=white
[go-url]: ./go.mod
[discord-badge]: https://img.shields.io/badge/Discord-join%20chat-5865F2?logo=discord&logoColor=white
[discord-url]: https://discord.gg/6THZ96357r
[x-badge]: https://img.shields.io/badge/X-%40RaylsLabs-000000?logo=x&logoColor=white
[x-url]: https://x.com/RaylsLabs
[linkedin-badge]: https://img.shields.io/badge/LinkedIn-Rayls-0A66C2?logo=linkedin&logoColor=white
[linkedin-url]: https://www.linkedin.com/company/rayls/
[youtube-badge]: https://img.shields.io/badge/YouTube-Rayls-FF0000?logo=youtube&logoColor=white
[youtube-url]: https://www.youtube.com/@Rayls_blockchain
