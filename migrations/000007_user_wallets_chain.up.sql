-- Add a chain discriminator to user_wallets so a user can hold private-chain and public-chain
-- wallets. Default 1 (WalletChainPrivate) preserves existing rows. The chain is a discriminator
-- only: it is NOT unique per (user_id, chain) — global wallet-address uniqueness (uq_rayls_address)
-- remains the only uniqueness invariant.
ALTER TABLE user_wallets ADD COLUMN chain INTEGER NOT NULL DEFAULT 1;
