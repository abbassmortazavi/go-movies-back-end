package repository

import (
	"backend/internal/models"
	"database/sql"
)

type Repository interface {
	Movies() ([]*models.Movie, error)
	Connection() *sql.DB
}
