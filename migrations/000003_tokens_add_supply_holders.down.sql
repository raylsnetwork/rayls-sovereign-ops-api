ALTER TABLE tokens
    DROP COLUMN IF EXISTS total_supply,
    DROP COLUMN IF EXISTS holder_count;
