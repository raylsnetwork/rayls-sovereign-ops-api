CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    email VARCHAR(255) NOT NULL DEFAULT '',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    status INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX idx_users_email ON users (email) WHERE email != '';

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

CREATE TABLE user_wallets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    rayls_address VARCHAR(42) NOT NULL,
    custody_provider INTEGER NOT NULL,
    custody_external_id VARCHAR(255) NOT NULL,
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

CREATE TABLE token_blacklist (
    id UUID PRIMARY KEY,
    jti VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_token_blacklist_expires_at ON token_blacklist (expires_at);
