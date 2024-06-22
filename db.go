package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

// Database connection (can be in-memory or a file).
var dbHandle *sql.DB

// Initialize database connection.
func initDB(dbPath, migrationSourceURL string) error {
	var err error
	dbHandle, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Couldn't get DB handle with path %s", dbPath)

	}
	driver, err := sqlite3.WithInstance(dbHandle, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Can't get DB driver for migrations: %s", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationSourceURL,
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("Can't get migrate instance: %s", err)
	}
	version, _, err := m.Version()
	if err != nil {
		log.Printf("Can't get DB version! %s", err)
	}
	log.Println("DB version is", version)
	err = m.Migrate(2)
	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("No migrations to run")
	} else if err != nil {
		log.Fatalf("Can't run migrations: %s", err)
	}
	return nil
}
