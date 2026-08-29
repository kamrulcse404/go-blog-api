package storage

import (
	"blogapi/models"
	"context"
	"errors"

	"github.com/lib/pq"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

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

		var pqErr *pq.Error

		if errors.As(err, &pqErr) && string(pqErr.Code) == "23505" && pqErr.Constraint == "users_email_key" {
			return models.User{}, ErrEmailAlreadyExists
		}

		return models.User{}, err
	}

	return user, nil
}
