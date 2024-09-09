package main

import (
	"social-network/internal/repository/sqlite"
)

func main() {
	sqlite.ApplyMigrations()
}
