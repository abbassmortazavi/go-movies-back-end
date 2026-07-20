package pkg

import (
	"backend/internal/repository/dbrepo"
	"database/sql"
)

func GetDB() *sql.DB {
	return dbrepo.DB
}
