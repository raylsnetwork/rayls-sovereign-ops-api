-- blockscout_balances_trigger.sql
--
-- Target:  Blockscout PostgreSQL database (NOT the ops-api database).
-- Apply:   Installed automatically by ops-api on startup via
--          infrastructure/database/blockscout.go:ApplyBlockscoutBalancesTrigger.
--          Manual: psql "$BLOCKSCOUT_DB_CONN" -f scripts/blockscout_balances_trigger.sql
--
-- Effect:  Every INSERT into address_current_token_balances and every UPDATE of
--          the `value` column fires pg_notify('balance_change', <json payload>).
--          ops-api listens on this channel via LISTEN and processes the payload
--          in cmd/api/indexer/blockscout_balances_listener.go.
--
-- Column names (address_hash, token_contract_address_hash, value, block_number,
-- id) match the Blockscout standard schema.
--
-- DELETE events are intentionally not captured: balance rows are not removed.
-- UPDATE is constrained to the `value` column so cosmetic writes (e.g.
-- value_fetched_at) do not generate notifications.
--
-- pg_notify has an 8000-byte payload limit. The JSON payload is small in
-- practice; values for ERC20 fit comfortably. If a notification is silently
-- dropped because of size, the next backfill sweep will pick the change up.
--
-- This script is idempotent: CREATE OR REPLACE and DROP TRIGGER IF EXISTS
-- make it safe to re-run without manual cleanup.

CREATE OR REPLACE FUNCTION notify_balance_change() RETURNS trigger AS $$
BEGIN
  PERFORM pg_notify('balance_change', json_build_object(
    'op',                          TG_OP,
    'address_hash',                encode(NEW.address_hash, 'hex'),
    'token_contract_address_hash', encode(NEW.token_contract_address_hash, 'hex'),
    'value',                       NEW.value::text,
    'block_number',                NEW.block_number
  )::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS on_balance_change ON address_current_token_balances;

CREATE TRIGGER on_balance_change
AFTER INSERT OR UPDATE OF value ON address_current_token_balances
FOR EACH ROW EXECUTE FUNCTION notify_balance_change();
