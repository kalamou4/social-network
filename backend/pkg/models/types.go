package models

import (
	"database/sql"
)

type Database struct {
	Db *sql.DB
}
