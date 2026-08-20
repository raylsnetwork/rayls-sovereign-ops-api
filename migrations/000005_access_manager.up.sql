CREATE TABLE IF NOT EXISTS am_roles (
    role_id          BIGINT      PRIMARY KEY,
    name             TEXT        NOT NULL DEFAULT '',
    label            TEXT        NOT NULL DEFAULT '',
    admin_role_id    BIGINT,
    guardian_role_id BIGINT,
    grant_delay_secs INT         NOT NULL DEFAULT 0,
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS am_role_members (
    role_id         BIGINT      NOT NULL,
    account         VARCHAR(42) NOT NULL,
    execution_delay INT         NOT NULL DEFAULT 0,
    active_since    TIMESTAMPTZ,
    is_active       BOOLEAN     NOT NULL DEFAULT true,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (role_id, account)
);

CREATE INDEX IF NOT EXISTS idx_am_role_members_account ON am_role_members (account);
CREATE INDEX IF NOT EXISTS idx_am_role_members_active  ON am_role_members (is_active) WHERE is_active = true;

CREATE TABLE IF NOT EXISTS am_managed_contracts (
    contract_address VARCHAR(42) PRIMARY KEY,
    deployer         VARCHAR(42) NOT NULL,
    is_paused        BOOLEAN     NOT NULL DEFAULT false,
    registered_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS am_function_permissions (
    contract_address VARCHAR(42) NOT NULL,
    selector         VARCHAR(10) NOT NULL,
    role_id          BIGINT      NOT NULL,
    is_active        BOOLEAN     NOT NULL DEFAULT true,
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (contract_address, selector, role_id)
);

CREATE INDEX IF NOT EXISTS idx_am_fn_perms_contract ON am_function_permissions (contract_address);
CREATE INDEX IF NOT EXISTS idx_am_fn_perms_role     ON am_function_permissions (role_id);

CREATE TABLE IF NOT EXISTS am_scheduled_operations (
    operation_id     VARCHAR(66) PRIMARY KEY,
    caller           VARCHAR(42) NOT NULL,
    managed_contract VARCHAR(42) NOT NULL,
    execute_after    TIMESTAMPTZ,
    status           VARCHAR(20) NOT NULL DEFAULT 'scheduled',
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_am_sched_ops_status   ON am_scheduled_operations (status);
CREATE INDEX IF NOT EXISTS idx_am_sched_ops_contract ON am_scheduled_operations (managed_contract);

CREATE TABLE IF NOT EXISTS am_contract_scoped_role_members (
    role_id          BIGINT      NOT NULL,
    account          VARCHAR(42) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    execution_delay  INT         NOT NULL DEFAULT 0,
    active_since     TIMESTAMPTZ,
    is_active        BOOLEAN     NOT NULL DEFAULT true,
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (role_id, account, contract_address)
);

CREATE INDEX IF NOT EXISTS idx_am_cs_members_account  ON am_contract_scoped_role_members (account);
CREATE INDEX IF NOT EXISTS idx_am_cs_members_contract ON am_contract_scoped_role_members (contract_address);

CREATE TABLE IF NOT EXISTS am_event_log (
    id               UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    event_name       TEXT        NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    block_number     BIGINT      NOT NULL,
    tx_hash          VARCHAR(66) NOT NULL,
    log_index        BIGINT      NOT NULL,
    block_time       TIMESTAMPTZ NOT NULL,
    payload          JSONB       NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (tx_hash, log_index)
);

CREATE INDEX IF NOT EXISTS idx_am_event_log_block   ON am_event_log (block_number);
CREATE INDEX IF NOT EXISTS idx_am_event_log_event   ON am_event_log (event_name);
CREATE INDEX IF NOT EXISTS idx_am_event_log_payload ON am_event_log USING gin (payload);
