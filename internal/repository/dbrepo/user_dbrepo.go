package dbrepo

import (
	"backend/internal/models"
	"context"
)

func (p *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select * users from users where email = $1`
	rows, err := p.DB.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user models.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
