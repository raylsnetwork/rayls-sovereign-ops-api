CREATE TABLE wallet_balances (
    id UUID PRIMARY KEY,
    wallet_address VARCHAR(42) NOT NULL,
    token_address VARCHAR(42) NOT NULL,
    balance TEXT NOT NULL,
    block_number BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX uq_wallet_balances_wallet_token ON wallet_balances (wallet_address, token_address);
CREATE INDEX idx_wallet_balances_wallet_address ON wallet_balances (wallet_address);
CREATE INDEX idx_wallet_balances_token_address ON wallet_balances (token_address);
