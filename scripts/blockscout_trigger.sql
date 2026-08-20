-- blockscout_trigger.sql
--
-- Target:  Blockscout PostgreSQL database (NOT the ops-api database).
-- Apply:   Run once by whoever manages the Blockscout deployment.
--          Example: psql "$BLOCKSCOUT_DB_CONN" -f scripts/blockscout_trigger.sql
--
-- Effect:  Every INSERT into tokens and every UPDATE of name, symbol, type,
--          total_supply, or holder_count fires pg_notify('token_change', <json payload>).
--          `type` is included so a re-classification by the RaylsTokenDiscovery fetcher
--          (e.g. Rayls-ERC-20 -> Rayls-StableCoin) propagates to ops-api's erc_standard even
--          when no other watched column changed.
--          ops-api listens on this channel via LISTEN and processes the
--          payload in cmd/api/indexer/blockscout_listener.go.
--
-- Column names (contract_address_hash, name, symbol, type, decimals,
-- total_supply, holder_count) were confirmed against the Blockscout DB schema.
-- `decimals` is included so the frontend can render supply in human units
-- (without it, decimals defaults to 0 and total_supply shows as raw base units).
--
-- DELETE events are intentionally not captured: tokens are never deleted from
-- the Blockscout tokens table once indexed.
--
-- pg_notify has an 8000-byte payload limit. The JSON payload is small in practice
-- (address + name + symbol + type + total_supply + holder_count), but if total_supply
-- grows unexpectedly large, PostgreSQL will silently drop the notification. Monitor
-- for missed events if tokens with extremely large supply values are expected.
--
-- This script is idempotent: CREATE OR REPLACE and DROP TRIGGER IF EXISTS
-- make it safe to re-run without manual cleanup.

CREATE OR REPLACE FUNCTION notify_token_change() RETURNS trigger AS $$
BEGIN
  PERFORM pg_notify('token_change', json_build_object(
    'op',           TG_OP,
    'address',      encode(NEW.contract_address_hash, 'hex'),
    'name',         NEW.name,
    'symbol',       NEW.symbol,
    'type',         NEW.type,
    'decimals',     NEW.decimals,
    'total_supply', NEW.total_supply::text,
    'holder_count', NEW.holder_count
  )::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS on_token_change ON tokens;

CREATE TRIGGER on_token_change
AFTER INSERT OR UPDATE OF name, symbol, type, decimals, total_supply, holder_count ON tokens
FOR EACH ROW EXECUTE FUNCTION notify_token_change();
