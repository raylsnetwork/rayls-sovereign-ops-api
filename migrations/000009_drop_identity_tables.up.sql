-- Accounts move to the shared identity database (ops_identity, migrations-identity).
--
-- A custody wallet is an EVM keypair, so ONE wallet per user works on every chain — only its
-- on-chain state (balance, roles) differs. Keeping users/wallets here meant each chain minted
-- its own wallet for the same person, so roles landed on one address while the token claimed
-- another. This database now holds chain-scoped facts only (am_*, tokens, balances, cursors).
--
-- The per-chain ops-api reads accounts over IDENTITY_DB_CONN. Set it before applying this, or
-- the instance has no user tables to fall back on.
DROP TABLE IF EXISTS user_providers;
DROP TABLE IF EXISTS user_wallets;
DROP TABLE IF EXISTS token_blacklist;
DROP TABLE IF EXISTS nonces;
DROP TABLE IF EXISTS users;
