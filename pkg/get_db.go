package pkg

import (
	"database/sql"
)

var DB *sql.DB

func GetDB() *sql.DB {
	return DB
}
