CREATE TABLE tokens (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name TEXT,
    symbol TEXT,
    resource_id TEXT CONSTRAINT chk_tokens_resource_id_nonempty CHECK (resource_id IS NULL OR resource_id != ''),
    metadata_url TEXT,
    erc_standard SMALLINT,
    decimals SMALLINT,
    issuer_id TEXT,
    status SMALLINT,
    contract_address VARCHAR(42) NOT NULL,
    token_class VARCHAR(50) NOT NULL DEFAULT 'unknown'
);
CREATE UNIQUE INDEX idx_tokens_contract_address ON tokens (contract_address);
CREATE UNIQUE INDEX idx_tokens_resource_id ON tokens (resource_id) WHERE resource_id IS NOT NULL;
CREATE INDEX idx_tokens_token_class ON tokens (token_class);


CREATE TABLE indexer_state (
    key VARCHAR(100) PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE token_events (
    id UUID PRIMARY KEY,
    contract_address VARCHAR(42) NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    block_number BIGINT NOT NULL,
    tx_hash VARCHAR(66) NOT NULL,
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_token_events_contract ON token_events (contract_address);
CREATE INDEX idx_token_events_block ON token_events (block_number);
CREATE INDEX idx_token_events_tx_hash ON token_events (tx_hash);
