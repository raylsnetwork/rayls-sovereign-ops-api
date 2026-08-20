-- Shared identity service schema.
--
-- This is the ONE database that spans every chain: who a person is, how they log in, and
-- which chains' ops-apis they may then talk to. Chain-scoped data (tokens, am_*, balances,
-- HSM signing wallets) lives in each chain's own ops-api database and never appears here.
--
-- Derived from the ops-api schema (migrations/000001_init.up.sql + 000008), keeping only
-- the identity tables. Column definitions are kept byte-identical so the same GORM models
-- and repositories serve both services.

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    email VARCHAR(255) NOT NULL DEFAULT '',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    status INTEGER NOT NULL DEFAULT 1,
    -- keccak256(user.id) — how the on-chain contracts know this user. A pure function of
    -- the UUID (domain.OnChainUserID), so it is the SAME on every chain, which is exactly
    -- why it belongs to identity rather than to any one chain. Persisted because keccak256
    -- cannot be recomputed in Postgres, and the per-chain services need the reverse map
    -- (bytes32 -> UUID) when listing pending address pairs read off-chain.
    on_chain_user_id BYTEA,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX idx_users_email ON users (email) WHERE email != '';
CREATE UNIQUE INDEX idx_users_on_chain_user_id ON users (on_chain_user_id)
    WHERE on_chain_user_id IS NOT NULL;

CREATE TABLE user_providers (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    provider INTEGER NOT NULL,
    oauth_id VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    wallet_address VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_oauth_provider UNIQUE (provider, oauth_id)
);
CREATE INDEX idx_user_providers_user_id ON user_providers (user_id);

-- Profile/lead answers collected by the standalone email sign-up form (company, headcount,
-- referral source, goals). Only the email path collects them — Google, Microsoft and SIWE
-- supply identity, not a questionnaire — so this is a 1:1 side table rather than four
-- always-empty columns on `users`, whose definition is kept byte-identical to the per-chain
-- schema so one GORM model serves both services.
CREATE TABLE user_signup_details (
    id UUID PRIMARY KEY,
    -- 1:1 with users: a person fills the form once. Re-submitting updates this row.
    user_id UUID NOT NULL UNIQUE REFERENCES users(id),
    company VARCHAR(255) NOT NULL DEFAULT '',
    -- Free-form bucket label as rendered by the form ("1-50", "51-200", ...), not a count.
    -- Stored as text so re-labelling the UI options cannot corrupt existing rows.
    employees VARCHAR(64) NOT NULL DEFAULT '',
    heard_about VARCHAR(255) NOT NULL DEFAULT '',
    goals TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_user_signup_details_user_id ON user_signup_details (user_id);

-- LOGIN wallets only — self-custody (MetaMask) addresses a user signs in with via SIWE.
-- These are credentials: chain-agnostic, one per person, never used to sign transactions.
--
-- HSM signing wallets are deliberately NOT here. They are minted per chain, funded there
-- and granted roles there, so an HSM wallet is meaningless away from its chain; they live
-- in the per-chain ops-api database. The custody_provider column keeps that distinction
-- explicit and lets the same model serve both sides.
CREATE TABLE user_wallets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    rayls_address VARCHAR(42) NOT NULL,
    custody_provider INTEGER NOT NULL,
    custody_external_id VARCHAR(255) NOT NULL,
    chain INTEGER NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX uq_rayls_address ON user_wallets (LOWER(rayls_address));
CREATE INDEX idx_user_wallets_user_id ON user_wallets (user_id);

CREATE TABLE nonces (
    id UUID PRIMARY KEY,
    wallet_address VARCHAR(42) NOT NULL,
    nonce VARCHAR(64) NOT NULL,
    message TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_nonces_wallet_address ON nonces (wallet_address);
CREATE UNIQUE INDEX idx_nonces_nonce ON nonces (nonce);
CREATE INDEX idx_nonces_expires_at ON nonces (expires_at) WHERE used = FALSE;

-- Revoked refresh-token jtis (NOT token contracts — see cmd/api/auth/token.go). Refresh
-- tokens are minted here, so revocation is checked here.
CREATE TABLE token_blacklist (
    id UUID PRIMARY KEY,
    jti VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_token_blacklist_expires_at ON token_blacklist (expires_at);
