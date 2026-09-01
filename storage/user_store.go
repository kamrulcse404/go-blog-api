package storage

import (
	"blogapi/models"
	"context"
	"errors"
	"fmt"

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

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `
		SELECT 
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE email=$1
	`
	var user models.User
	err := DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return models.User{}, fmt.Errorf("get user by email: %w", err)
	}

	return user, nil
}

func GetUserByID(ctx context.Context, id int) (models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE id=$1
	`

	var user models.User

	err := DB.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}
