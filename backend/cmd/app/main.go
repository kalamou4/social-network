package main

import (
	"social-network/backend/db/sqlite"
)

func main() {

	sqlite.ApplyMigrations()

}
