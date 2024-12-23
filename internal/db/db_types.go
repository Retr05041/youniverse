package db

import "database/sql"

type Database struct {
	instance *sql.DB
}
