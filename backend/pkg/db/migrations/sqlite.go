package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase(dbpath string) (*Database, error) {
	Db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, fmt.Errorf("failed to oepn database : %w", err)
	}

	return &Database{Db: Db}, nil
}

func (d *Database) Migrate() error {
	
	driver, err := sqlite3.WithInstance(d.Db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not create driver : %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		os.Getenv("MIGRATION_PATH"),
		"sqlite3", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil

}
