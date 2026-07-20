package repository

import (
	"backend/internal/models"
)

type Repository interface {
	Movies() ([]*models.Movie, error)
}
