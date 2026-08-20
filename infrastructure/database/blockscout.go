package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
	"github.com/raylsnetwork/rayls-privacy-ops-api/scripts"
)

// ApplyBlockscoutTrigger connects to the Blockscout database and installs the
// notify_token_change trigger. Idempotent and safe to call on every startup.
// Skips with a warning if connStr is empty.
func ApplyBlockscoutTrigger(ctx context.Context, connStr string, log logger.Logger) error {
	if connStr == "" {
		log.Warn("BLOCKSCOUT_DB_CONN not set — skipping trigger installation")
		return nil
	}

	db, err := sql.Open("postgres", normalizeSSLMode(connStr))
	if err != nil {
		return fmt.Errorf("open blockscout db: %w", err)
	}
	defer func() { _ = db.Close() }()

	if _, err = db.ExecContext(ctx, scripts.BlockscoutTriggerSQL); err != nil {
		return fmt.Errorf("apply blockscout trigger: %w", err)
	}

	log.Info("Blockscout trigger applied")
	return nil
}

// ApplyBlockscoutBalancesTrigger installs the notify_balance_change trigger on
// Blockscout's address_current_token_balances table. Idempotent and safe to
// call on every startup. Skips with a warning if connStr is empty.
func ApplyBlockscoutBalancesTrigger(ctx context.Context, connStr string, log logger.Logger) error {
	if connStr == "" {
		log.Warn("BLOCKSCOUT_DB_CONN not set — skipping balances trigger installation")
		return nil
	}

	db, err := sql.Open("postgres", normalizeSSLMode(connStr))
	if err != nil {
		return fmt.Errorf("open blockscout db: %w", err)
	}
	defer func() { _ = db.Close() }()

	if _, err = db.ExecContext(ctx, scripts.BlockscoutBalancesTriggerSQL); err != nil {
		return fmt.Errorf("apply blockscout balances trigger: %w", err)
	}

	log.Info("Blockscout balances trigger applied")
	return nil
}
