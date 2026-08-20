# Architecture

How the Rayls ops platform is put together: what the services are, why they are split the
way they are, and where each piece of state lives.

For the request-by-request walkthroughs (login, chain creation, token deploy), see
[flows.md](./flows.md).

---

## The shape of the system

There is no single "ops-api". There are **two kinds** of service, and the split is the most
important thing to understand:

| | Identity service | ops-api |
|---|---|---|
| **How many** | One, shared | One **per chain** |
| **Answers** | "Who is this person?" | "What may they do *here*, and do it" |
| **Owns** | users, login methods, sessions | tokens, balances, roles, indexers |
| **Knows about chains** | Nothing. No RPC, no contracts | Exactly one chain |
| **Database** | `ops_identity` | `ops_api_<slug>`, one per chain |
| **Lifecycle** | Long-lived | Created and destroyed **with its chain** |

Both are the same binary (`rayls-ops-api`), started with different subcommands.

```
                    ┌──────────────────┐
   browser ────────>│    playground    │  (Next.js, :3000)
                    └────────┬─────────┘
                             │
              /auth/*        │        /api/tokens, /api/wallets …
         ┌───────────────────┴───────────────────┐
         v                                       v
┌──────────────────┐                  ┌────────────────────────┐
│ identity service │                  │  ops-api for chain A   │
│      :8090       │   identity JWT   │  + worker + custody    │
│                  │ ················>│  + its own database    │
│  ops_identity DB │  (shared secret) └────────────────────────┘
└──────────────────┘                  ┌────────────────────────┐
         │                  ········> │  ops-api for chain B   │
         │                                 … one per chain
         v
    ┌─────────┐
    │ custody │  (HSM — chain-agnostic, shared)
    └─────────┘
```

### Why per-chain

A single ops-api serving every chain was the source of a class of bugs: delete a chain,
create a new one, and the same process with the same database kept serving the **old**
chain's tokens, while its cached registry address made new deploys fail.

Because a chain's ops-api and its database are now created and destroyed *with the chain*,
that failure is not prevented — it is impossible. There is no shared state to go stale.

### Why identity is not per-chain

A person is the same person on every chain they own. If each chain owned its own `users`
table, one human would be several unrelated accounts with several unrelated wallets, and
"one account, many chains" could never work.

So identity is extracted, and the seam is clean: **no chain-scoped table has a foreign key
into the identity tables**, and every database transaction stays within one side of the
split.

---

## The three entrypoints

```bash
rayls-ops-api run       --config <env>   # per-chain HTTP API        (:8080)
rayls-ops-api worker    --config <env>   # per-chain indexers
rayls-ops-api identity  --config <env>   # shared identity service   (:8090)
```

`run` and `worker` are per chain and share that chain's database. `identity` is the single
shared service. All three build the same binary; `run`/`worker` use `di.New`, `identity`
uses `di.NewIdentity` — a deliberately smaller graph with no chain dependencies at all.

---

## Where state lives

### Identity database (`ops_identity`, `migrations-identity/`)

| Table | Holds |
|---|---|
| `users` | the person; `on_chain_user_id` = keccak256(uuid), identical on every chain |
| `user_providers` | how they log in (Google, Microsoft, email, SIWE) |
| `user_wallets` | **login** wallets only — see the custody split below |
| `user_signup_details` | profile answers from the email sign-up form (company, headcount, referral, goals) |
| `nonces` | SIWE challenges |
| `token_blacklist` | revoked refresh-token JTIs (not token contracts) |

`user_signup_details` is 1:1 with `users` and exists **only for the email path** — Google,
Microsoft and SIWE supply identity, not a questionnaire, so they never write a row. It is
kept out of `users` because those columns are neither identity nor credentials, and because
the `users` row is deliberately byte-identical to the per-chain schema (one GORM model
serves both). Re-submitting the form updates the row rather than inserting a second.

### Per-chain database (`ops_api_<slug>`, `migrations/`)

| Table | Holds |
|---|---|
| `tokens`, `token_events` | tokens on this chain |
| `wallet_balances` | balances on this chain |
| `am_*` (7 tables) | this chain's AccessManager state, kept current by the indexer |
| `indexer_state` | indexer cursors |

This database holds **chain-scoped facts only**. Accounts are not here: migration `000009`
dropped `users`, `user_providers`, `user_wallets`, `nonces` and `token_blacklist` from the
per-chain schema, and the per-chain container reads them over `IDENTITY_DB_CONN` instead.

A custody wallet is an EVM keypair, so **one wallet works on every chain** — only its
on-chain state (balance, roles) differs. Minting one per chain meant each chain minted a
*different* address for the same person, so roles landed on one address while the token
claimed another. One wallet, shared, is what makes `user_wallets` an identity table.

`user_wallets.chain` is therefore **not** a chain-instance discriminator. It separates the
private-chain wallet from the public-chain one (`domain.WalletChain`: `Private=1`,
`Public=2`), and is deliberately not unique per `(user_id, chain)` — a user accumulates
both. Global address uniqueness (`uq_rayls_address`) is the only invariant.

`custody_provider` still distinguishes what a row *means*:

- **`Self`** — the user's own MetaMask address. A *credential*: SIWE identifies them by it.
- **`RaylsHSM`** — a minted signing key, usable on any chain the user reaches.

---

## Authentication

### One token, many chains

The identity service mints the session JWT. Every chain's ops-api verifies it using a
**shared `JWT_SECRET`** — chains do not mint tokens and have no login endpoints of their own.

The token deliberately carries **identity only**:

| Claim | Present | Why |
|---|---|---|
| `user.id`, `name`, `email` | yes | who they are — true everywhere |
| `auth_method` | yes | how they logged in |
| `is_admin` | when email matches `OPS_ADMIN_EMAIL` | deployment-wide |
| `roles` | **no** | a grant on chain A says nothing about chain B |
| `custody_wallet_address` | **no** | wallets are minted per chain |

`NewIdentityTokenWrapper` enforces this: it takes no wallet repository, and it strips
`roles` *even if the caller passes them*. That last part matters — the chain-less login
path returns a placeholder operator role to keep dev instances usable, and it must never
travel across chains as an authorization claim.

Access tokens live 15 minutes; refresh tokens 24 hours and rotate on use.

### Authorization is per chain, per request

`RequireRole` asks that chain's `ChainRoleService`, which resolves:

```
user id (from JWT) → their HSM wallet on THIS chain → am_role_members → role names
```

All local reads — the AccessManager indexer has already mirrored the chain state, so no
RPC is on the request path. Two consequences worth knowing:

- A **revoked role takes effect immediately**, rather than lingering until the user's next
  login. Roles used to be a login-time snapshot frozen into the JWT.
- A token minted for one chain **cannot** authorize another, because the claim is not
  consulted at all when a resolver is wired.

When the AccessManager is not wired (a chain-less instance), the resolver is nil and
`RequireRole` falls back to the JWT claim — preserving single-service behaviour exactly.

### First contact with a chain

A user has **one** custody wallet, held in the identity database and valid on every chain.
What they lack on a chain they have never used is not the wallet but its on-chain state:
no roles there, so no access.

`ProvisioningService.Provision` runs on every login (`auth_service.go`). It returns the
existing wallet if there is one and mints via custody only when there is none, then
`prepareWalletForChain` funds it for gas and grants `FACTORY_DEPLOYER` +
`PRIVACY_NODE_OPERATOR` on *this* chain.

Funding and granting are best-effort: both can be done out of band, and neither should
fail the request that triggered them. The mint is **not** idempotent — see Known gaps.

### Custody

The HSM is told neither a chain id nor an RPC URL — it just holds keys and signs bytes.
That is why one custody service is shared and one minted key works on every chain. The
binding between a user and a wallet exists only in `user_wallets`; the HSM does not know
about users, and `POST /api/wallet` has no field that could carry a user id.

**Signing is always delegated to custody.** The key never leaves the HSM.

### OAuth across many hosts

Providers only redirect to pre-registered URIs, but chains are created on demand, so
per-chain callback URLs cannot be registered in advance.

Every service therefore sets `OAUTH_REDIRECT_BASE` to the **playground origin**, and carries
its own callback URL inside the OAuth `state`. One registered URI
(`http://localhost:3000/auth/google/callback` in dev) serves everything; the playground
relays the callback to whichever service began the flow.

That relay validates the bounce target against the ops URLs RayUp actually reports, and
**fails closed**. The state is attacker-controllable and an authorization code is a bearer
credential, so a path check alone would not be enough.

---

## Provisioning (RayUp)

RayUp is the control plane. On chain creation it stands up, in order: the chain itself →
its explorer → contracts → **its ops stack** (ops-api + worker + custody, one Helm release,
its own databases) → its playground. `Instance.opsUrl` records where that chain's ops-api
lives; the playground routes data requests there.

`destroyOps` reverses it, dropping `ops_api_<slug>` and `raylzdb_<slug>` with the chain.

**Cloud vs local** differ only at the edge. Cloud gives each ops-api an ingress host with
TLS. Local k3d has neither cert-manager nor wildcard DNS, so `dev.enabled=true` swaps the
ingress for a NodePort that k3d publishes on the host (`http://localhost:<port>`), and
points the chart at the Postgres and NATS already running in compose instead of standing
up in-cluster copies. Same chart, same per-chain isolation.

Local NodePort allocation: chain RPC ports are handed out **upward** from the bottom of the
k3d-mapped band; ops ports are mirrored **downward** from the top. They cannot collide
until the band is exhausted, and the code errors rather than silently overlapping.

---

## Project layout

```
cmd/api/
  app/            entrypoints: api.go (run), worker.go, identity.go, cmd.go (cobra)
  core/           domain errors + port interfaces
  adapters/
    handlers/     HTTP handlers (Swagger-annotated)
    repositories/ GORM repositories
    blockchain/   on-chain services (AccessManager, factory, tokens, governance)
    custody/      Rayls HSM adapter (+ self-custody no-op)
  services/       domain services (auth, provisioning, roles, permissions, wallets)
  auth/           SIWE / OAuth / JWT (go-pkgz/auth)
  indexer/        Blockscout + AccessManager indexers (run in `worker`)
  messaging/      NATS/JetStream
  middleware/     auth, role resolution, query validation
  sse/            in-memory SSE hub
di/               container.go (per-chain) + identity_container.go (shared)
domain/           domain models
config/           Viper configuration
migrations/           per-chain schema
migrations-identity/  identity schema
contracts/        abigen v2 bindings
infrastructure/   database connection + migrations
```

Dependency wiring happens in `di/`, never in `app.go` — the entrypoints pull handlers from
a container and register routes.

---

## Configuration

Shared across services:

| Variable | Meaning |
|---|---|
| `JWT_SECRET` | **must be identical** on identity and every ops-api |
| `DATABASE_CONNECTIONSTRING` | that service's own database |
| `OAUTH_REDIRECT_BASE` | the playground origin (the OAuth relay) |
| `CORSURLS` | `;`-separated allowed origins |

Identity only:

| Variable | Default | Meaning |
|---|---|---|
| `IDENTITY_PORT` | `8090` | listen port |
| `BOOTSTRAP_TOKEN` | *(empty)* | guards `/admin/bootstrap`; empty = unauthenticated |
| `EMAIL_SIGNUP_ENABLED` | `false` | **dev only** — see Security |

Per-chain ops-api:

| Variable | Meaning |
|---|---|
| `BLOCKCHAIN_RPC_URL`, `DEPLOYMENT_PROXY_REGISTRY_ADDR` | this chain; both set ⇒ role resolver active |
| `BLOCKSCOUT_DB_CONN`, `NATS_URL`, `INSTANCE_NAME` | indexers, messaging, queue isolation |
| `CHAINLESS` | boot with no chain bound — relaxes config validation and self-disables every on-chain feature (dev; the chain arrives later and the stack is rebound, see flows.md) |

---

## Security notes

**`/admin/bootstrap` creates the first admin** and is otherwise unauthenticated — guarded
only by "no users exist yet". On a fresh deployment the first caller owns it. Set
`BOOTSTRAP_TOKEN` (checked in constant time) or keep the endpoint unreachable from outside.
Now that identity is shared, that admin is the admin of *every* chain.

**`POST /auth/signup` is off by default.** There is no verification code yet, so it issues
a valid session to whoever submits an address — including the admin's. `EMAIL_SIGNUP_ENABLED`
must stay `false` outside a local machine. The same gap makes everything it stores in
`user_signup_details` self-reported and unverified, down to the email the row hangs off:
treat those answers as lead information, never as an attested fact about a customer.

**Other measures:** SIWE nonces are single-use (`SELECT FOR UPDATE`, 5-minute TTL); refresh
tokens rotate and are blacklisted by JTI on use and logout; session cookies are `HttpOnly`,
with `Secure` on wherever TLS terminates; the `?from=` login parameter is validated against
`CORSURLS` to prevent open redirects.

---

## Known gaps

- **Custody wallet creation is still not idempotent, but orphans are now traceable.**
  `POST /api/wallet` takes only `type`, `password` and `address_quantity` — no user id,
  no idempotency key — and generates a fresh random key per call, with no delete API. That
  cannot be fixed from this side; closing it fully needs an idempotency key on the custody
  endpoint.
  What *is* fixed: every mint is now **write-ahead** (`cmd/api/services/wallet_mint.go`).
  An intent row is committed before the HSM call, so the database is authoritative before
  the irreversible side effect. A crash mid-mint leaves a row naming the user instead of an
  untraceable key, and the retry no longer mints a fresh orphan every attempt.
  - The intent row reuses `is_active=false` with a `pending:` placeholder in
    `rayls_address`, so it is invisible to every existing query — **no migration needed**.
  - A failed mint deletes its intent (no key exists, nothing leaks). Only a crash between
    minting and completion still orphans a key, and that now leaves a durable pointer to it;
    `ensureWalletFor` logs stranded intents on the user's next pass.
  - Reattaching an orphan automatically is still impossible: the only HSM lookup is
    `GET /api/wallet/address/{address}`, which needs the address that was lost.
- **Chain status reconciliation lives in RayUp, and is now in place.** A row left
  `PROVISIONING` by a worker that died mid-create used to stay there forever: the only path
  that sets `ERROR` is the job's own `catch`, which never runs if the process is killed.
  `apps/worker/src/reconciler.ts` now re-reads Kubernetes on an interval and corrects
  stranded rows (`PROVISIONING`/`RESTARTING` → `RUNNING` or `ERROR` by pod health,
  interrupted `DESTROYING` → `DELETED`, `QUEUED` with no job → `ERROR`). It runs in the
  worker because that process is a singleton (`replicas: 1` + `Recreate`), so it needs no
  leader election. Staleness is measured from `DeployJob.startedAt`, never
  `Instance.updatedAt`, which every progress tick refreshes.
