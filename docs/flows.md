# Flows

Request-by-request walkthroughs of the Rayls ops platform. For the structural picture —
what the services are and why they are split — see [architecture.md](./architecture.md).

Throughout: **identity** is the one shared service on `:8090`; **ops-api** is a per-chain
service (`:8080` locally, one per chain in the cluster).

---

## Endpoint map

### Identity service (`rayls-ops-api identity`, `:8090`)

| Method | Path | Auth | Notes |
|---|---|---|---|
| GET | `/health` | — | |
| POST | `/admin/bootstrap` | `BOOTSTRAP_TOKEN` if set | creates the first admin; one-shot |
| GET | `/auth/google/login` · `/auth/microsoft/login` | — | starts OAuth |
| GET | `/auth/google/callback` · `/auth/microsoft/callback` | — | reached via the playground relay |
| GET | `/auth/siwe/login?address=` | — | SIWE challenge |
| POST | `/auth/siwe/callback` | — | verify signature, issue session |
| POST | `/auth/signup` | — | **off unless `EMAIL_SIGNUP_ENABLED=true`**; stores the form's profile answers |
| POST | `/auth/refresh` | refresh token | rotates the pair |
| POST | `/auth/logout` | — | blacklists the refresh token, clears the cookie |
| GET | `/auth/user` · `/auth/status` · `/auth/list` | session | go-pkgz built-ins |
| GET | `/api/me` | session | **not** role-gated — this service has no chain |

### ops-api (`rayls-ops-api run`, per chain)

Authenticated = valid identity JWT. Operator = additionally holds
`PRIVACY_NODE_OPERATOR` **on this chain**.

| Method | Path | Auth |
|---|---|---|
| GET | `/health` · `/swagger/*` | — |
| ANY | `/auth/*` | — | present, but the playground routes auth to identity |
| GET | `/api/me` | operator |
| POST · GET | `/api/me/address-pairs` | authenticated |
| GET | `/api/tokens` · `/api/tokens/:address` | operator |
| POST | `/api/tokens` · `/api/tokens/estimate` | authenticated |
| GET | `/api/tokens/:address/permissions` | authenticated |
| POST | `/api/tokens/:address/mint` · `/burn` · `/teleport` | authenticated (+ on-chain permission) |
| POST · GET | `/api/tokens/:address/register` · `/api/tokens/registry[/pending]` | authenticated |
| GET | `/api/tokens/stream` | authenticated (SSE) |
| GET | `/api/wallets/:address/balances[/:tokenAddress]` | authenticated |
| GET | `/api/wallets/balances/stream` | authenticated (SSE) |
| GET · PATCH | `/api/v1/admin/*` | operator |

Mint/burn/teleport are *authenticated* at the route and then checked against the chain's
AccessManager inside the handler — the role gate alone is not what authorizes them.

---

## Login

All methods end the same way: identity sets an `HttpOnly` JWT cookie on the **playground
origin**, so it is first-party to the app and travels with later API calls.

### SIWE (MetaMask)

```
browser                     playground              identity
   │                            │                      │
   ├─ GET /auth/siwe/login?address=0x… ───────────────>│  rewrite → :8090
   │                            │                      ├─ generate nonce (32 bytes)
   │                            │                      ├─ build EIP-4361 message
   │                            │                      └─ store nonce (5 min, unused)
   │<────────────────── { message, nonce } ────────────┤
   │
   ├─ personal_sign in MetaMask
   │
   ├─ POST /auth/siwe/callback { address, signature, nonce } ─────>│
   │                            │                      ├─ consume nonce (SELECT FOR UPDATE,
   │                            │                      │    single-use: rejects replay)
   │                            │                      ├─ ecrecover → compare to claimed address
   │                            │                      ├─ first time? create user +
   │                            │                      │    self-custody wallet + SIWE link
   │                            │                      └─ login decision tree (below)
   │<──── Set-Cookie: JWT + { refresh_token } ─────────┤
```

New wallets are **auto-registered** on first sight. There is no separate registration step
and no registration token.

### Google / Microsoft

```
browser              playground                identity              provider
   │                     │                        │                     │
   ├─ /auth/google/login ─────────────────────────>│                     │
   │                     │                        ├─ state = { csrf, own callback URL }
   │                     │                        └─ redirect ──────────>│
   │<─────────────────────── consent ────────────────────────────────────┤
   │                     │                        │                     │
   ├─ callback (code, state) ──> playground        │                     │
   │                     ├─ decode state.r
   │                     ├─ check target is a KNOWN ops/identity origin
   │                     │    (fails closed — the code is a bearer credential)
   │                     └─ 307 → identity/auth/google/callback ────────>│
   │                                              ├─ exchange code, fetch profile
   │                                              ├─ find-or-create user
   │                                              └─ login decision tree
   │<──────── Set-Cookie: JWT + #refresh_token ───┤
```

One registered redirect URI serves every service — see *OAuth across many hosts* in
architecture.md.

### One email, two sign-in methods

An email that already belongs to an account is only linked to a *new* provider when the
account has **no** provider yet — the pre-created bootstrap admin. Otherwise the login is
refused with `409 EMAIL_ALREADY_LINKED`, naming the provider the account does use.

This is not an approval problem, and the account may be fully approved: the email sign-up
path asserts `emailVerified` without verifying anything, so linking on a bare address match
would hand an established account to whoever types its email. Signing up with an address
already used for Google therefore fails by design — sign in with Google instead.

Lifting this needs a real verification step (see the TODO on `EmailSignup`), not a looser
check here.

### The login decision tree

Applied by `applyLoginDecisionTree` after any method identifies the user:

| Condition | Result |
|---|---|
| `is_active = false` | `403 ACCOUNT_SUSPENDED` |
| `status = waiting_role_assignment`, provisioner wired | activate the account, continue |
| `status = waiting_role_assignment`, no provisioner | `403 ROLE_ASSIGNMENT_PENDING` |
| chain-less instance | return placeholder role — **dropped when an identity token is minted** |
| otherwise | read roles from the chain's AccessManager |
| … RPC failed | `503 SERVICE_UNAVAILABLE` |
| … no roles | `403 ROLE_ASSIGNMENT_PENDING` |

In the **identity service** the last three rows never apply: it has no chain, so it
activates the account and mints an identity-only token. Authorization is deferred to
whichever chain the user then talks to.

This replaces the earlier SDD login-flow documents (removed), which described
`waiting_role_assignment` as a hard rejection and predated both auto-provisioning and the
identity split.

### What a login writes

Every method converges on `FindOrCreateOAuthUser`, so the account-creation writes are the
same; email is modelled as a provider whose `oauth_id` *is* the address. All of it lands in
the identity database.

| Table | Google / Microsoft | Email sign-up | SIWE |
|---|---|---|---|
| `users` | `name` from the provider | `name` = the email's local part | name blank until set |
| `user_providers` | `oauth_id` = provider `sub`/`id` | `oauth_id` = the email | `wallet_address` |
| `user_wallets` | minted HSM wallet | minted HSM wallet | the user's own address (`Self`) |
| `user_signup_details` | — | company, employees, heardAbout, goals | — |

Written on **first** login only. Afterwards it is a lookup by `(provider, oauth_id)`: the
wallet already exists and the status is already `role_assigned`, so nothing is written —
except a repeat email sign-up, which overwrites `user_signup_details` in place.

Two deliberate properties of the details write: it is **best-effort** (a failure is logged
and the session still issues — these are lead data, not credentials), and a request
carrying none of the four fields **skips the write entirely**, so an email-only API call
cannot blank out answers already on file.

### Admin bootstrap (once per deployment)

```
POST /admin/bootstrap { "email": "admin@example.com" }
  Authorization: Bearer <BOOTSTRAP_TOKEN>        # required when configured

  ├─ refuse if any user exists                 → 409
  ├─ mint an HSM custody wallet
  ├─ create users + user_wallets + provider link   (one transaction)
  └─ 201 { "address": "0x…" }                  → grant it on-chain roles per chain
```

---

## Authenticated request to a chain

```
browser ──> playground ──> ops-api for the selected chain
   cookie      proxy adds       │
                the cookie      ├─ RequireAuth: verify JWT with the SHARED secret
                                │    (identity signed it; this service did not)
                                │
                                ├─ RequireRole → ChainRoleService:
                                │    user id → their HSM wallet ON THIS CHAIN
                                │      └─ none? mint + fund + grant now (first contact)
                                │    wallet → am_role_members → role names
                                │    (local reads — the indexer already mirrored the chain)
                                │
                                └─ handler
```

Outcomes worth distinguishing:

| Response | Meaning |
|---|---|
| `401` | no session, or the signature did not verify |
| `403` | authenticated, but holds no matching role **on this chain** |
| `200` | authorized here |

A `403` where you expected `200` usually means the user's wallet on this chain has not been
granted its roles yet — not that login failed.

---

## Creating a chain

```
playground "Create Sovereign Chain"
   └─> RayUp control plane
         ├─ allocate slug, chainId, RPC NodePort
         ├─ install the chain (genesis → validators → RPC health)
         ├─ per-chain explorer (Blockscout)
         ├─ deploy contracts (one-off in-cluster Job)
         ├─ ops stack for THIS chain:
         │     ├─ register business roles on its AccessManager
         │     ├─ helm install ops-<slug>  (ops-api + worker + custody)
         │     ├─ its own databases: ops_api_<slug>, raylzdb_<slug>
         │     ├─ wait healthy → bootstrap its admin → grant on-chain roles
         │     └─ record Instance.opsUrl
         └─ status RUNNING (only after ALL of the above)
```

`RUNNING` means *fully set up*, not merely "the chain answers RPC" — the status flips only
after contracts and the ops stack are in place.

The playground then routes that chain's data requests to its `opsUrl`. Auth still goes to
identity.

---

## Deleting a chain

```
RayUp destroy
   ├─ helm uninstall ops-<slug>
   ├─ DROP DATABASE ops_api_<slug>, raylzdb_<slug>
   ├─ remove the per-chain explorer + its database
   └─ delete the chain namespace
```

**This is why a new chain cannot show a previous chain's tokens.** Its ops-api and database
are gone; the next chain gets its own, empty. The identity database is untouched — accounts
outlive chains.

---

## Token deploy, mint, burn

```
POST /api/tokens                       (authenticated)
  ├─ resolve caller → their HSM wallet on this chain
  ├─ build calldata for RNContractFactory
  ├─ custody signs (the key never leaves the HSM)
  ├─ broadcast, poll for the receipt
  └─ store the token, status Internal, issuerId = this chain

POST /api/tokens/:address/mint | /burn  (authenticated)
  ├─ resolve caller → their HSM wallet on this chain
  ├─ check the AccessManager permission for THIS function on THIS token
  │     └─ missing → 403
  ├─ build calldata per ERC standard
  └─ custody signs → broadcast → receipt
```

Deploy is open to any authenticated user; mint and burn additionally require the on-chain
permission, checked in the handler rather than by the route's role gate.

---

## How the token list stays current

```
chain ──> Blockscout ──> ops-api worker ──────> ops_api_<slug>.tokens
                             │                          │
                             ├─ JetStream (durable)     │
                             └─ core NATS ──> ops-api ──┴──> SSE ──> browser
                                              (hub fans out)
```

Two paths on purpose: JetStream carries durable events, core NATS carries the live event the
API fans out over SSE. A cursor-based backfiller plus `LISTEN/NOTIFY` keeps the projection
current. Wallet balances follow the same shape, filtered to addresses present in
`user_wallets` for this chain.

---

## Local development

```bash
./start_dev.sh --rayup            # keeps cluster, registry, images and caches — fast restart
./start_dev.sh --rayup --clean    # destroys EVERYTHING, rebuilds from scratch — slow
```

`--clean` runs in two phases, and the order matters: chains and their explorers are
destroyed **while the control plane is still up** (destroying a chain is a job its worker
runs, and the explorer databases live in the shared Postgres), and only then is the
infrastructure itself — ops stack, control plane, shared Blockscout, k3d cluster *and its
registry*, images, volumes — torn down. Everything is then rebuilt.

Skipping straight to the hard teardown would orphan RayUp's port allocations and strand
`blockscout_rayup-*` databases in a Postgres that comes back empty.

### Why ops-api runs before any chain exists

Under `--rayup` there is no chain at boot — chains are created from the playground, on
demand — so the host ops-api and worker start with `CHAINLESS=true`:

- config validation stops requiring `PN_RPC_URL` / `BLOCKCHAIN_RPC_URL` /
  `DEPLOYMENT_PROXY_REGISTRY_ADDR`;
- every on-chain feature self-disables — no role resolver, no indexers, no deploy;
- the HTTP surface still serves: health, swagger, auth pass-through, `/api/me/address-pairs`;
- login takes the chain-less branch and returns a placeholder operator role, so the instance
  stays usable while it has no AccessManager to ask.

The point is that `localhost:8080` is a stable address the playground can hold from the
first second, with a warm process behind it. It stays chain-less for the whole run: chains
never bind to it, because each one brings its own ops stack.

### How the chain gets wired in

No user action beyond creating the chain. RayUp provisions each chain **with its own
ops-api, worker and Blockscout** (helm releases `ops-<slug>` / `bs-<slug>`, each on its own
database) and destroys them with it. Nothing is rebound, because nothing is shared.

```
playground "Create Sovereign Chain"
   └─> RayUp: install chain → deploy contracts
        ├─ register business roles on that chain's AccessManager
        ├─ grant FACTORY_DEPLOYER + PRIVACY_NODE_OPERATOR to the creator + admin wallets
        ├─ helm install ops-<slug>  (ops-api + worker, own database)
        ├─ helm install bs-<slug>   (that chain's Blockscout)
        └─ status RUNNING

playground request naming a chain
   └─> routed to THAT chain's ops-api (src/lib/chain-target.ts),
       ownership-checked against the caller
```

The host ops-api on `localhost:8080` is not in that path — it serves the chain-less surface
only (health, swagger, auth pass-through, `/api/me/address-pairs`).

Each chain has its OWN ops database, so indexer cursors and the `am_*` mirror can never
carry another chain's state. Accounts are not in there: users and their one custody wallet
live in the shared `ops_identity`, so a chain can be destroyed without taking its users
with it.

This replaced a model that bound ONE shared ops-api to ONE chain at a time. The second user
to create a chain never got the slot, so their chain was running and deployed but unusable,
and their requests were served by the first user's stack.

What comes up locally:

| Service | Port |
|---|---|
| playground | 3000 |
| ops-api (host, chain-less for the whole run) | 8080 |
| **identity** | **8090** |
| custody | 5032 |
| per-chain ops-api in k3d | NodePort, from the top of the mapped band |
| per-chain explorer | allocated per instance |

The single registered OAuth callback is `http://localhost:3000/auth/google/callback`.

### Verifying the split by hand

```bash
# 1. identity mints an identity-only token
curl -s -c jar -X POST localhost:8090/auth/signup \
     -H 'Content-Type: application/json' \
     -d '{"email":"you@example.com","name":"You",
          "company":"Acme Bank","employees":"51-200",
          "heardAbout":"Conference","goals":"Issue a tokenised deposit."}'
# needs EMAIL_SIGNUP_ENABLED=true. Everything after "name" is optional and lands in
# user_signup_details; omit it all and only users/user_providers/user_wallets are written.

# 2. the claims carry no roles and no wallet
curl -s -b jar localhost:8090/auth/user
# {"…","attrs":{"auth_method":"email"}}

# 3. a chain's ops-api verifies a token it did not mint
curl -s -b jar localhost:8080/auth/user        # → the same user

# 4. …but authorizes independently
curl -s -o /dev/null -w '%{http_code}\n' -b jar localhost:8080/api/tokens   # 403 without roles here
curl -s -o /dev/null -w '%{http_code}\n'        localhost:8080/api/tokens   # 401 with no session
```

Step 3 returning the user while step 4 returns `403` is the whole design in two requests:
one login, honoured everywhere; authorization decided per chain.

---

## Troubleshooting

| Symptom | Cause |
|---|---|
| `401` everywhere after login | `JWT_SECRET` differs between identity and ops-api |
| `403` on `/api/tokens`, login fine | no roles on **this** chain — wallet not granted yet |
| New chain lists an old chain's tokens | a shared ops-api/database — per-chain provisioning did not run |
| Deploy fails with a stale registry | same cause: the ops-api is bound to a different chain |
| OAuth returns "not a known ops-api" | the relay allowlist could not reach RayUp, or `opsUrl` is unset — it fails closed |
| `/auth/signup` returns 404 | `EMAIL_SIGNUP_ENABLED` is false (the default) |
| Bootstrap returns 401 | `BOOTSTRAP_TOKEN` set; send `Authorization: Bearer <token>` |
| Bootstrap returns 409 | already bootstrapped — one admin per deployment |
| Roles unchanged after an on-chain grant | the AccessManager indexer has not caught up; roles come from `am_*` |
