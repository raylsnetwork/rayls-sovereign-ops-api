package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/withstack"
)

// RunMigrations applies the per-chain ops-api schema (the "migrations" folder).
func RunMigrations(connStr string) error {
	return RunMigrationsFrom(connStr, "migrations")
}

// RunMigrationsFrom applies the migrations in `folder` (a directory name resolved upward
// from the working directory, as the binary may run from anywhere).
//
// Two schemas live in this repo and they are deliberately separate databases:
//   - "migrations"          the per-chain ops-api (tokens, am_*, balances, HSM wallets)
//   - "migrations-identity" the shared identity service (users, providers, nonces)
//
// They are never applied to the same database: golang-migrate keeps a single
// schema_migrations table per database, so two version sequences would fight.
func RunMigrationsFrom(connStr, folder string) error {
	connStr = normalizeSSLMode(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return withstack.Wrap(err)
	}
	defer func() { _ = db.Close() }()

	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{})
	if err != nil {
		return withstack.Wrap(err)
	}

	migrationsPath, err := findProjectMigrationsPath(folder)
	if err != nil {
		return withstack.Wrap(err)
	}

	logger.Info("Running migrations", "path", migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver,
	)
	if err != nil {
		return withstack.Wrap(err)
	}
	defer func() { _, _ = m.Close() }()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("All migrations already applied")
			return nil
		}
		return withstack.Wrap(err)
	}

	logger.Info("All pending migrations applied successfully")

	return nil
}

func normalizeSSLMode(connStr string) string {
	connStr = strings.ReplaceAll(connStr, "sslmode=enable", "")
	if !strings.Contains(connStr, "sslmode=") {
		connStr += " sslmode=disable"
	}
	return connStr
}

func findProjectMigrationsPath(folder string) (string, error) {
	startDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := startDir
	for {
		migrationsPath := filepath.Join(dir, folder)
		if _, err := os.Stat(migrationsPath); err == nil {
			return migrationsPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find %s folder from: %s", folder, startDir)
		}
		dir = parent
	}
}
