package storage

import (
	"blogapi/models"
	"context"
)

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (
			name,
			email, 
			password_hash
		)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, created_at, updated_at
	`

	err := DB.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
