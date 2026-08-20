# Frontend Guide — Wallet Balances API

How the frontend integrates with the ops-api **wallet balances** endpoints. Balances are sourced from
Blockscout's `address_current_token_balances` table via two cooperating indexers (a cursor backfill at
startup + a `LISTEN/NOTIFY` listener), and surfaced to the client as a REST endpoint plus a live SSE
stream.

## General conventions

- **Base:** all endpoints below live under `/api/*`.
- **Authentication:** session via the `JWT` cookie (go-pkgz/auth). Every call must send
  `credentials: 'include'`. No valid session → `401`.
- **Addresses:** the backend normalizes to lowercase. You may send checksum or lowercase; compare
  case-insensitively in the UI.
- **Coverage:** balances are only tracked for wallets present in `user_wallets`. Querying any other
  address returns `404`.
- **Availability:** the SSE stream is only registered when NATS is configured; otherwise it 404s.

| Method | Route | Auth | What it does |
|---|---|---|---|
| GET | `/api/wallets/{address}/balances` | authenticated | Lists current per-token balances for a wallet |
| GET | `/api/wallets/{address}/balances/{tokenAddress}` | authenticated | Returns one balance, filtered by wallet and token |
| GET | `/api/wallets/balances/stream` | authenticated | SSE stream of balance updates for the caller's wallet |

---

## 1) List balances — `GET /api/wallets/{address}/balances`

Returns the current balances for the given wallet, joined with token metadata.

### Response — `200 OK`

```json
[
  {
    "walletAddress": "0xaaaa...aaaa",
    "tokenAddress": "0xbbbb...bbbb",
    "tokenSymbol": "TKN",
    "tokenName": "Test Token",
    "decimals": 18,
    "balance": "1000000000000000000",
    "blockNumber": 12345,
    "updatedAt": "2026-06-09T12:00:00.000000Z"
  }
]
```

- `balance` is the raw on-chain amount as a decimal **string** (full uint256 precision). Use a big-int
  library on the frontend; **do not** parse it as a JS `Number`.
- `tokenSymbol` / `tokenName` / `decimals` may be empty when the token is brand new and the tokens
  indexer hasn't picked it up yet. Render gracefully.

### Errors

| Status | When |
|---|---|
| `401` | No session cookie |
| `404` | The address is not in `user_wallets` |

---

## 1b) Get one balance — `GET /api/wallets/{address}/balances/{tokenAddress}`

Returns the balance for a single (wallet, token) pair. Same shape as one entry of the list endpoint,
not wrapped in an array. Useful when the UI already knows which token it's rendering and wants a
direct refresh instead of refetching the whole list.

### Response — `200 OK`

```json
{
  "walletAddress": "0xaaaa...aaaa",
  "tokenAddress": "0xbbbb...bbbb",
  "tokenSymbol": "TKN",
  "tokenName": "Test Token",
  "decimals": 18,
  "balance": "1000000000000000000",
  "blockNumber": 12345,
  "updatedAt": "2026-06-09T12:00:00.000000Z"
}
```

### Errors

| Status | When |
|---|---|
| `401` | No session cookie |
| `404` | The wallet is not in `user_wallets`, **or** the wallet has no stored balance for the token |

The two 404 cases are distinguishable by the response body: the wallet-not-found case carries
`wallet not found`; the missing-pair case carries the `wallet balance` resource identifier with the
`wallet/token` pair in the message.

---

## 2) Live stream — `GET /api/wallets/balances/stream`

Server-Sent Events stream of balance updates, scoped to the **authenticated user's own wallet**. The
backend looks up the caller's `user_wallets.rayls_address` and filters the broadcast to events whose
`walletAddress` matches.

```js
const es = new EventSource('/api/wallets/balances/stream', { withCredentials: true });
es.addEventListener('wallet_balance', (e) => {
  const evt = JSON.parse(e.data);
  // evt = { type: 'balance_updated', walletAddress, tokenAddress, balance, blockNumber }
});
```

- Heartbeat: a `: ping` comment line is emitted every 25 seconds to keep proxies from idling out.
- Each delivered event mirrors the durable JetStream payload — same JSON, no transformation.
- Closing the tab unregisters the client from the in-memory hub.

### Event payload

```json
{
  "type": "balance_updated",
  "walletAddress": "0xaaaa...aaaa",
  "tokenAddress": "0xbbbb...bbbb",
  "balance": "1000000000000000000",
  "blockNumber": 12345
}
```

`type` is currently always `balance_updated`. If we add new event types in the future they will
appear here without changing the SSE event name.
