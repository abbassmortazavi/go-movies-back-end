package repository

import (
	"backend/internal/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}
