# Frontend Authentication Guide

This document describes how to integrate with the ops-api authentication system. The frontend uses WAGMI for wallet interactions.

## API Base URL

Call every `/auth/*` endpoint below **same-origin on the playground** (`http://localhost:3000`
in dev). The playground proxies them to the shared identity service; keeping them
same-origin is what makes the session cookie first-party, so it is sent on later `/api/*`
calls.

Note `/auth/*` and `/api/*` are served by *different* services: auth by the one shared
identity service, data by the ops-api of whichever chain is selected. The playground routes
both — see [architecture.md](./architecture.md).

## Error Format

All error responses use the same format:

```json
{ "error": "description of what went wrong" }
```

## Checking Auth Status

To check if the user is logged in (e.g., on page load):

```
GET /auth/user
```

- **Authenticated**: returns `200` with user info (`name`, `id`, `email`, `attributes`)
- **Not authenticated**: returns `401`

The browser automatically sends the `JWT` cookie — no manual token handling needed.

The `attributes` object includes:

| Key | Type | Notes |
|---|---|---|
| `auth_method` | string | `"siwe"` \| `"google"` \| `"microsoft"` \| `"refresh"` |
| `roles` | string[] | On-chain roles from RaylsAccessManager (e.g. `["BANK_EMPLOYEE"]`). Omitted when empty. |
| `custody_wallet_address` | string | The user's wallet address as held by their custody provider (HSM or self-custody), normalized lowercase. Omitted when the user has no wallet yet. |

The `custody_wallet_address` is embedded at JWT mint time, so it is refreshed on every login and every `/auth/refresh`. If the wallet is provisioned after login, the frontend must call `/auth/refresh` to pick it up.

## Google / Microsoft OAuth

Simplest flow — just redirect the user:

1. Navigate to `GET /auth/google/login` (or `/auth/microsoft/login`)
2. User completes consent on Google/Microsoft
3. The server sets a `JWT` cookie and redirects back to the root URL (`/`)
4. Done — all subsequent requests automatically include the cookie

The frontend should handle the redirect landing (e.g., check auth status on `/` and route accordingly).

## Standalone email sign-up

`POST /auth/signup` authenticates by email alone and returns the same session as any other flow. It is **disabled by default** and returns `404` unless the backend sets `EMAIL_SIGNUP_ENABLED=true` — there is no verification code yet, so it must not be reachable outside a local machine. Treat a `404` here as "not enabled", not as a bug.

```json
POST /auth/signup
{
  "email": "you@example.com",
  "name": "You",
  "company": "Acme Bank",
  "employees": "51-200",
  "heardAbout": "Conference",
  "goals": "Issue a tokenised deposit."
}
```

Only `email` is required. `name` defaults to the local part of the address. The remaining four fields are the sign-up form's profile answers and are stored in the identity database against the user; send them together or omit them together.

Two behaviours worth designing around:

- **The profile answers never block the session.** If storing them fails the login still succeeds, so a `200` does not by itself confirm they were written.
- **Omitting all four skips the write.** An email-only re-signup leaves previously submitted answers intact rather than blanking them; sending them again overwrites in place.

Response matches the other flows — a `JWT` cookie plus `refresh_token` in the body. Because the answers are self-reported and the address is unverified, do not surface them anywhere they would read as an attested fact about a customer.

## SIWE (MetaMask) via WAGMI

Three-step flow using WAGMI hooks for wallet interaction.

### Step 1 — Connect wallet and get challenge

Use WAGMI's `useAccount` and `useConnect` to get the wallet address, then fetch the challenge from the API:

```ts
const { address } = useAccount();

const res = await fetch(`/auth/siwe/login?address=${address}`, {
  credentials: 'include',
});
const { message, nonce } = await res.json();
```

### Step 2 — Sign with WAGMI and verify

Use WAGMI's `useSignMessage` to sign the challenge:

```ts
const { signMessageAsync } = useSignMessage();

const signature = await signMessageAsync({ message });

const res = await fetch('/auth/siwe/callback', {
  method: 'POST',
  credentials: 'include',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ address, signature, nonce }),
});
const data = await res.json();
```

Two possible responses:

- Existing user (cookie is set):
  ```json
  {
    "status": "authenticated",
    "user": { "id": "...", "name": "...", "role": "..." }
  }
  ```
There is **no separate registration step**. A wallet signing in for the first time is
registered automatically and receives a session in the same response.

## After Login (All Flows)

- Auth is cookie-based — the browser sends the `JWT` cookie automatically
- Protected endpoints live under `/api/*` — they return **401** if no valid cookie
- Role-restricted endpoints return **403** when the user holds no matching role

**401 and 403 mean different things.** 401 is "not logged in". 403 is "logged in, but you
hold no role *on this chain*" — roles are per chain, resolved on each request against that
chain's AccessManager. See [architecture.md](./architecture.md) and [flows.md](./flows.md).

To renew a session explicitly:

```
POST /auth/refresh    { "refresh_token": "..." }
```

Returns a rotated `refresh_token` and sets a fresh `JWT` cookie. The old refresh token is
revoked, so store the new one.

## Logout

```
POST /auth/logout     { "refresh_token": "..." }
```

**Must be POST.** A GET falls through to a different handler and the refresh token is never
revoked. Sending the refresh token blacklists it; omitting it still clears the cookie.

```ts
const { disconnect } = useDisconnect();

async function logout(refreshToken: string) {
  await fetch('/auth/logout', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ refresh_token: refreshToken }),
  });
  disconnect();
}
```

## CORS

Prefer same-origin paths (`/auth/user`, not `http://localhost:8090/auth/user`) — the
playground proxies them, so the cookie stays first-party. Either way, always pass
`credentials: 'include'`:

```ts
fetch('/auth/user', { credentials: 'include' });
```

Without it the browser will neither send nor accept the cookie. If you do call a service
directly on its own origin, that origin must appear in the service's `CORSURLS`.

## Gotchas

- SIWE verify uses a **POST with JSON body** (`{ address, signature, nonce }`), not query params
- The `refresh_token` is returned in the **JSON response**, not as a cookie — hold it in memory and send it to `/auth/refresh` and `/auth/logout`
- Refresh tokens **rotate**: each refresh invalidates the previous one, so always store the token you just received
- All wallet addresses are lowercased server-side — the frontend can send any case
- The cookie name is `JWT` — useful for debugging in browser DevTools (Application > Cookies)
- WAGMI's `signMessageAsync` returns the signature in `0x...` hex format, which is what the API expects
