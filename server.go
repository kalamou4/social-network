package main

import (
	"log"
	"os"
	migrations "social-network/backend/pkg/db/migrations"
	config "social-network/backend/config"

	_ "github.com/mattn/go-sqlite3"

)

func main() {

	err := config.LoadEnvFile(".env")
	if err != nil {
		return
	}

	db, err := migrations.NewDatabase(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Rest of your application logic
}
