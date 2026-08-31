package storage

import (
	"blogapi/models"
	"context"
	"database/sql"
)

func GetPosts(ctx context.Context, limit int, offset int) ([]models.Post, error) {

	query := `SELECT id, user_id, title, content,  created_at, updated_at
		FROM posts
		ORDER BY created_at DESC, id DESC
		LIMIT $1 OFFSET $2`

	rows, err := DB.QueryContext(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post

		err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	query := `
		INSERT INTO posts (user_id, title, content)
			VALUES ($1, $2, $3)
		RETURNING id, user_id, title, content, created_at, updated_at
	`

	err := DB.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetPostByID(ctx context.Context, id int) (models.Post, error) {
	query := `
		SELECT id, user_id, title, content, created_at, updated_at
		FROM posts 
		WHERE id = $1
	`
	var post models.Post

	err := DB.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func UpdatePost(ctx context.Context, id int, updatedPost models.Post) (models.Post, error) {

	query := `
		UPDATE posts
		SET
			title = $1,
			content = $2,
			updated_at = NOW()
		WHERE id = $3
		RETURNING id, title, content, created_at, updated_at
	`

	var post models.Post

	err := DB.QueryRowContext(
		ctx,
		query,
		updatedPost.Title,
		updatedPost.Content,
		id,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func DeletePost(ctx context.Context, id int) error {

	query := `
		DELETE FROM posts 
		WHERE id = $1
	`

	result, err := DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func CountPosts(ctx context.Context) (int, error) {
	var total int

	query := `
		SELECT COUNT(*)
		FROM posts
	`

	err := DB.QueryRowContext(ctx, query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
