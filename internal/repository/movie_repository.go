package repository

import (
	"backend/internal/models"
)

type MovieRepository interface {
	Movies() ([]*models.Movie, error)
}
