# Frontend Guide — Tokens API

How the frontend integrates with the ops-api **token** endpoints: deploy, query, permissions,
mint/burn, and the real-time stream (SSE).

## General conventions

- **Base:** all endpoints below live under `/api/*`.
- **Authentication:** session via the **`JWT` cookie** (go-pkgz/auth). Every call must send
  `credentials: 'include'`. No valid session → `401`.
- **CORS:** if the frontend runs on a different origin, it must be listed in `CORSURLS` (credentialed
  CORS is already enabled on the backend).
- **Addresses:** the backend normalizes to lowercase. You may send checksum or lowercase — when
  comparing on the frontend, use a case-insensitive comparison.
- **Availability:** mint/burn, permissions and SSE only exist when the backend has the Access Manager /
  NATS configured; otherwise the route is not registered (`404`).

| Method | Route | Auth | What it does |
|---|---|---|---|
| POST | `/api/tokens` | authenticated | Deploys a token via the factory |
| GET | `/api/tokens` | `PRIVACY_NODE_OPERATOR` | Lists tokens (paginated) |
| GET | `/api/tokens/{address}` | `PRIVACY_NODE_OPERATOR` | Token details |
| GET | `/api/tokens/{address}/permissions` | authenticated | What the user can do on the token |
| POST | `/api/tokens/{address}/mint` | authenticated | Mints tokens |
| POST | `/api/tokens/{address}/burn` | authenticated | Burns tokens |
| POST | `/api/tokens/{address}/pause` | the token's `pauser` | Pauses/resumes a stablecoin |
| GET | `/api/tokens/stream` | authenticated | SSE stream of token events |

### Enums

`ercStandard` (number) — also returned as `ercStandardLabel` (the protocol's canonical name):

| value | label |
|---|---|
| `0` Custom | `CUSTOM` |
| `1` ERC-20 | `RAYLS_ERC20` |
| `2` ERC-721 | `RAYLS_ERC721` |
| `3` ERC-1155 | `RAYLS_ERC1155` |
| `4` Enygma (= ERC-20) | `RAYLS_ENYGMA` |
| `5` ERC-721 DVP | `RAYLS_ERC721_DVP` |
| `6` ERC-1155 DVP | `RAYLS_ERC1155_DVP` |

`status` (number) — also returned as `statusLabel`:

| value | label | meaning |
|---|---|---|
| `0` | `inactive` | |
| `1` | `active` | live on-chain (e.g. organically discovered tokens) |
| `2` | `paused` | |
| `3` | `frozen` | |
| `4` | `deprecated` | |
| `10` | `internal` | deployed via the API; **stays internal** (not promoted to Active) |

> Prefer the `*Label` fields in the UI instead of hardcoding the numeric values.

---

## 1) Deploy a token — `POST /api/tokens`

Deploys **one token per request** via the factory, signed by the logged-in user's custody wallet.

**Body**

```json
{
  "standard": "ERC20",
  "name": "My Token",
  "symbol": "MTK",
  "decimals": 18,
  "uri": "ipfs://..."
}
```

| Field | Required | Notes |
|---|---|---|
| `standard` | yes | `ERC20`, `ENYGMA`, `ERC721`, `ERC1155`, `ERC721_DVP`, `ERC1155_DVP` (case-insensitive) |
| `name` / `symbol` | depends | ERC20/Enygma and ERC721 require both; ERC1155 requires `name` |
| `decimals` | optional | ERC20/Enygma |
| `uri` | depends | ERC721/ERC1155 |

> `resourceId` is **not** sent on deploy — the backend uses `bytes32(0)`.

**Prerequisites:** the user must have a custody wallet, and that wallet must hold the on-chain deploy
role (e.g. `FACTORY_DEPLOYER`) — otherwise the tx reverts (`422`).

**Response `201`**

```json
{ "deployedAddress": "0xabc...", "txHash": "0xdef...", "standard": "ERC20" }
```

> **Synchronous** call: it only responds after the on-chain receipt (a few seconds). The token is stored
> with `status: 10` (`internal`) and **stays internal** — it is not promoted to Active by the indexer.
> `issuerId` holds the **chainId** of the PN it was deployed on.

**Errors:** `400` (body/standard/params, or user without a wallet) · `401` · `422` (revert, e.g. missing
role) · `502` (chain failure).

```js
const res = await fetch('/api/tokens', {
  method: 'POST', credentials: 'include',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ standard: 'ERC20', name: 'My Token', symbol: 'MTK', decimals: 18 }),
});
const { deployedAddress, txHash } = await res.json();
```

---

## 2) List tokens — `GET /api/tokens`

Paginated list (role `PRIVACY_NODE_OPERATOR`). Query: `page`, `limit` (max 100), `name`, `symbol`, `type`.
Returns a `PagedResponse` whose items match the detail shape below.

---

## 3) Details — `GET /api/tokens/{address}`

Role `PRIVACY_NODE_OPERATOR`. `address` is case-insensitive.

**Response `200`**

```json
{
  "id": "uuid",
  "name": "My Token",
  "symbol": "MTK",
  "contractAddress": "0xabc...",
  "ercStandard": 1,
  "ercStandardLabel": "RAYLS_ERC20",
  "tokenClass": "unknown",
  "status": 10,
  "statusLabel": "internal",
  "totalSupply": "1000",
  "holderCount": 3,
  "decimals": 18,
  "metadataUrl": "ipfs://...",
  "issuerId": "1337"
}
```

- `status` (number) + `statusLabel` (ready string: `inactive`/`active`/`paused`/`frozen`/`deprecated`/`internal`). Use `statusLabel` in the UI.
- `ercStandard` (number) + `ercStandardLabel` (protocol name: `RAYLS_ERC20`, `RAYLS_ENYGMA`, …; `CUSTOM` for non-protocol tokens). Use `ercStandardLabel` in the UI.
- `issuerId`: chainId of the PN the token was deployed on (set on API deploy; omitted when empty).

**Errors:** `401` · `403` · `404` (`{"error":"token not found"}`) · `500`.

```js
const res = await fetch(`/api/tokens/${address}`, { credentials: 'include' });
if (res.status === 404) { /* not found */ }
const token = await res.json();
```

---

## 4) User permissions — `GET /api/tokens/{address}/permissions`

Returns, **for the logged-in user's wallet**, what it can do on the token (derived from the Access
Manager). Use it to enable/disable action buttons. Authenticated (any user; uses their own wallet).

**Response `200`**

```json
{
  "contractAddress": "0xabc...",
  "walletAddress": "0xdef...",
  "isPaused": false,
  "canMint": true,
  "canBurn": true,
  "canSubmitTokenUpdate": false,
  "canPause": false,
  "isTokenPaused": false,
  "functions": [
    { "selector": "0x40c10f19", "name": "mint" },
    { "selector": "0x9dc29fac", "name": "burn" }
  ]
}
```

- `functions`: **only what the user can call** (`{selector, name}`; `name` may be empty if unknown). If
  nothing is allowed (or the token is not AM-managed): `functions: []` and all flags `false`.
- `canMint`/`canBurn`/`canSubmitTokenUpdate`: convenience flags derived from the list.
- `canPause`/`isTokenPaused`: **stablecoin only**, and *not* derived from the AM — see the two-pauses
  note below. Both `false` for every other standard.

**Two different "paused" fields — do not confuse them:**

| Field | Means | Source |
|---|---|---|
| `isPaused` | the **Access Manager** paused this managed contract | `am_managed_contracts`, from indexed events |
| `isTokenPaused` | the **stablecoin's own** `paused` flag — what actually halts transfers | live contract read |

Drive a pause/unpause button off `isTokenPaused`, never `isPaused`. Likewise `canPause` is not an AM
role: `pause()` is `onlyPauser`, a `msg.sender` equality check against the contract's own `pauser`
address, so it is true only for that one wallet.

**Errors:** `401` · `400` (user without a wallet) · `500`.

```js
const perms = await (await fetch(`/api/tokens/${address}/permissions`, { credentials: 'include' })).json();
mintButton.disabled = !perms.canMint;
burnButton.disabled = !perms.canBurn;
```

---

## 5) Mint — `POST /api/tokens/{address}/mint`

Mints tokens, signed by the logged-in user's wallet. The backend only executes if the wallet holds the
mint permission in the AM (`403` otherwise). Check `canMint` first in the UI.

**Body (varies by the token's `ercStandard`)**

| Standard | Body |
|---|---|
| ERC-20 / Enygma | `{ "to": "0x..", "amount": "1.5" }` |
| ERC-721 (and DVP) | `{ "to": "0x..", "tokenId": "123" }` |
| ERC-1155 (and DVP) | `{ "to": "0x..", "tokenId": "1", "amount": "10", "data": "0x" }` |

## 6) Burn — `POST /api/tokens/{address}/burn`

Burns tokens (`canBurn` permission; `403` otherwise).

**Body**

| Standard | Body |
|---|---|
| ERC-20 / Enygma | `{ "from": "0x..", "amount": "1.5" }` |
| ERC-721 (and DVP) | `{ "tokenId": "123" }` |
| ERC-1155 (and DVP) | `{ "from": "0x..", "tokenId": "1", "amount": "10" }` |

### Field rules (mint and burn)

- `to` / `from`: EVM address (`0x...`).
- **`amount`**: a **human** decimal string (e.g. `"1.5"`). The backend **scales it by the token's
  decimals**. It cannot have more places than `decimals`, nor be negative.
- **`tokenId`**: a **raw** integer string (not scaled).
- `data` (ERC-1155, optional): hex `0x...`.
- **DVP** standards use an empty `extraData` automatically on mint.

**Response `200` (mint and burn)**

```json
{ "txHash": "0xabc..." }
```

> **Synchronous** call (waits for the receipt). Treat it as async in the UI.

**Errors (mint and burn):** `400` (params) · `401` · `403` (no permission) · `404` (token) ·
`422` (on-chain revert) · `502` (chain failure).

```js
// Mint ERC-20
await fetch(`/api/tokens/${address}/mint`, {
  method: 'POST', credentials: 'include',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ to: recipient, amount: '1.5' }),
});

// Burn ERC-721
await fetch(`/api/tokens/${address}/burn`, {
  method: 'POST', credentials: 'include',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ tokenId: '123' }),
});
```

> Even with `canMint/canBurn = true`, if the AM index is momentarily behind the tx may revert (`422`) —
> treat it as a recoverable error.

---

## 7) Pause / unpause — `POST /api/tokens/{address}/pause`

Halts or resumes **all transfers, mints and burns** on a stablecoin, signed by the logged-in user's
custody wallet. `RAYLS_STABLECOIN` only — every other standard returns `400`.

```json
{ "paused": true }
```

`paused` is **required**: an omitted field is `400`, never silently read as `false` (which would
unpause). Send `true` to pause, `false` to resume.

**Response `200`** — `{ "txHash": "0x..." }`

**Errors**

| Code | Meaning |
|---|---|
| `400` | not a stablecoin, or `paused` missing/malformed |
| `403` | the caller's wallet is not the contract's `pauser` |
| `409` | the token is already in the requested state (nothing to do) |
| `422` | reverted on-chain |

Authorization is **not** an Access Manager role — `pause()` is `onlyPauser`, a `msg.sender` check
against the contract's own `pauser` address. Drive the button off `canPause` and `isTokenPaused` from
the permissions endpoint (§4), not off `canMint`-style AM flags or the registry `status`.

```js
const perms = await (await fetch(`/api/tokens/${address}/permissions`, { credentials: 'include' })).json();
if (perms.canPause) {
  await fetch(`/api/tokens/${address}/pause`, {
    method: 'POST',
    credentials: 'include',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ paused: !perms.isTokenPaused }),
  });
}
```

---

## 8) Real-time stream (SSE) — `GET /api/tokens/stream`

Pushes token lifecycle events (no polling). Useful to react when a token's supply/holders change or when
it is discovered on-chain.

- Cookie auth → use `EventSource(url, { withCredentials: true })`.
- Broadcast: the client receives **all** token events; **filter by `address`** on the frontend.

Each message arrives as `event: token`:

```
event: token
data: {"type":"discovered","address":"0x...","name":"...","symbol":"...","ercStandard":"ERC-20","totalSupply":"1000","holderCount":3,"status":"internal"}
```

- `type`: `discovered` (token discovered) or `supply_updated` (supply/holders changed).
- `status`: the token's current status (e.g. `internal` for API-deployed tokens, `active` for organically
  discovered ones).
- There are heartbeats (`: ping`) that `EventSource` ignores — no need to handle them.

```js
const es = new EventSource('/api/tokens/stream', { withCredentials: true });

es.addEventListener('token', (e) => {
  const evt = JSON.parse(e.data); // { type, address, status, ... }
  if (evt.address.toLowerCase() === deployedAddress.toLowerCase()) {
    // token discovered/updated → refresh the UI or re-fetch GET /api/tokens/{address}
  }
});

es.onerror = () => { /* the browser reconnects on its own; re-fetch GET on reconnect to resync */ };
```

> **At-most-once:** events during a reconnect window may be lost. The DB/chain is the source of truth — on
> (re)connect, re-fetch `GET /api/tokens` to resync.

### Infra (for whoever hosts it)
- Behind nginx: `proxy_buffering off;` for the stream route (the backend already sends `X-Accel-Buffering: no`).
- Do not gzip `text/event-stream`.

---

## End-to-end flow (example)

1. **Deploy:** `POST /api/tokens` → keep `deployedAddress` (the token is `internal`; `issuerId` = the PN's chainId).
2. **Follow:** open the SSE stream; when a `token` event with your `address` arrives, refresh the UI (or
   re-fetch `GET /api/tokens/{address}`).
3. **Permissions:** `GET /api/tokens/{address}/permissions` → enable the mint/burn buttons.
4. **Action:** `POST /api/tokens/{address}/mint` or `/burn` with the body for the token's standard.
