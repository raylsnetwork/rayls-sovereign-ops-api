-- Add the on-chain user identity to users: keccak256(user.ID), as the contract knows the user via
-- getAllPendingAddressPairs(). keccak256 is one-way and cannot be recomputed in Postgres, so the hash
-- is persisted at onboarding time to allow reverse-mapping bytes32 -> UUID for admin discovery.
-- Nullable: populated for onboarded users only (no backfill). A partial unique index allows many NULLs
-- while keeping non-null hashes unique.
ALTER TABLE users ADD COLUMN on_chain_user_id BYTEA;
CREATE UNIQUE INDEX idx_users_on_chain_user_id ON users (on_chain_user_id)
    WHERE on_chain_user_id IS NOT NULL;
