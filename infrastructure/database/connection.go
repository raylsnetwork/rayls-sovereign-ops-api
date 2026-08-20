package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
	"github.com/raylsnetwork/rayls-privacy-ops-api/withstack"
)

// Connect opens the per-chain ops-api database, creating and migrating it if needed.
func Connect(connStr string) (*gorm.DB, error) {
	return ConnectWithMigrations(connStr, "migrations")
}

// ConnectWithMigrations is Connect with an explicit migration folder, so a service can
// apply ITS OWN schema. The identity service passes "migrations-identity": its database
// holds only users/providers/nonces, and golang-migrate tracks one version sequence per
// database — applying both sets to one database makes the second run fail on a version
// it cannot find.
func ConnectWithMigrations(connStr, migrationsFolder string) (*gorm.DB, error) {
	client, err := openDB(connStr)
	if err != nil && strings.Contains(err.Error(), "does not exist") {
		logger.Info("Database not found. Attempting to create it...")
		params := parseDSN(connStr)
		if createErr := createDatabase(params); createErr != nil {
			return nil, withstack.Wrap(createErr)
		}
		client, err = openDB(connStr)
		if err != nil {
			return nil, withstack.Wrap(err)
		}
	}
	if err != nil {
		return nil, withstack.Wrap(err)
	}

	if err := RunMigrationsFrom(connStr, migrationsFolder); err != nil {
		return nil, withstack.Wrap(err)
	}

	logger.Info("Connected to the database successfully")
	return client, nil
}

// Open connects to an EXISTING database without creating or migrating it.
//
// For databases another service owns: the per-chain ops-api reads accounts from the shared
// identity database, whose schema the identity service applies (migrations-identity).
// golang-migrate keeps one version sequence per database, so a second migrator running the
// per-chain set against it would fight — and a reader has no business creating it either:
// if it is missing, that is a deployment error worth surfacing, not something to paper over.
func Open(connStr string) (*gorm.DB, error) {
	client, err := openDB(connStr)
	if err != nil {
		return nil, withstack.Wrap(err)
	}
	return client, nil
}

func Disconnect(client *gorm.DB) error {
	db, err := client.DB()
	if err != nil {
		return withstack.Wrap(err)
	}
	return db.Close()
}

func openDB(connStr string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
}

func parseDSN(dsn string) map[string]string {
	params := make(map[string]string)

	if strings.HasPrefix(dsn, "jdbc:postgresql://") {
		parseJdbcUri(dsn, params)
	} else {
		parseGormUri(dsn, params)
	}

	return params
}

func parseJdbcUri(dsn string, params map[string]string) map[string]string {
	dsn = strings.TrimPrefix(dsn, "jdbc:postgresql://")

	hostAndRest := strings.SplitN(dsn, "/", 2)
	if len(hostAndRest) == 2 {
		hostPort := hostAndRest[0]
		dbname := hostAndRest[1]

		params["dbname"] = dbname

		if strings.Contains(hostPort, ":") {
			hostParts := strings.SplitN(hostPort, ":", 2)
			params["host"] = hostParts[0]
			params["port"] = hostParts[1]
		} else {
			params["host"] = hostPort
		}
	}

	return params
}

func parseGormUri(dsn string, params map[string]string) map[string]string {
	parts := strings.Fields(dsn)
	for _, p := range parts {
		tokens := strings.SplitN(p, "=", 2)
		if len(tokens) == 2 {
			params[tokens[0]] = tokens[1]
		}
	}

	return params
}

func createDatabase(params map[string]string) error {
	required := []string{"host", "port", "user", "password", "dbname"}
	for _, key := range required {
		if _, ok := params[key]; !ok {
			return fmt.Errorf("missing required DSN param: %s", key)
		}
	}

	baseDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		params["host"], params["port"], params["user"], params["password"])

	createDBTimeout := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(createDBTimeout)*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", baseDSN)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	if pingErr := db.PingContext(ctx); pingErr != nil {
		return pingErr
	}

	query := fmt.Sprintf("CREATE DATABASE %s", params["dbname"])
	_, err = db.Exec(query)
	return err
}
