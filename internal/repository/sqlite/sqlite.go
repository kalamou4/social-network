package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"social-network/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase(dbpath string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, fmt.Errorf("failed to oepn database : %w", err)
	}
	return DB, nil
}

func Migrate(d *sql.DB) error {

	driver, err := sqlite3.WithInstance(d, &sqlite3.Config{})
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

func ApplyMigrations() {

	err := utils.LoadEnvFile(".env")
	if err != nil {
		return
	}

	db, err := ConnectDatabase(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = Migrate(db)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Rest of your application logic
}
