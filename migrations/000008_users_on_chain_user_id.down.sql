DROP INDEX IF EXISTS idx_users_on_chain_user_id;
ALTER TABLE users DROP COLUMN on_chain_user_id;
