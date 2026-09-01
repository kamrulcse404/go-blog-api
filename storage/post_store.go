package storage

import (
	"blogapi/models"
	"context"
	"database/sql"
)

func GetPosts(ctx context.Context, limit int, offset int, search string) ([]models.Post, error) {

	query := `SELECT 
				posts.id,
				posts.user_id, 
				posts.title, 
				posts.content,  
				posts.created_at, 
				posts.updated_at,

				users.id,
				users.name,
				users.email

			FROM posts
			JOIN users 
				ON posts.user_id=users.id

			WHERE posts.title ILIKE $1
			OR posts.content ILIKE $1
		
			ORDER BY posts.created_at DESC, posts.id DESC
			LIMIT $2 OFFSET $3
		`

	searchPattern := "%" + search + "%"

	rows, err := DB.QueryContext(ctx, query, searchPattern, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		var author models.Author

		err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,

			&author.ID,
			&author.Name,
			&author.Email,
		)

		if err != nil {
			return nil, err
		}

		post.Author = &author

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

func UpdatePost(ctx context.Context, id int, userID int, updatedPost models.Post) (models.Post, error) {

	query := `
		UPDATE posts
		SET
			title = $1,
			content = $2,
			updated_at = NOW()
		WHERE id = $3
		AND user_id = $4
		RETURNING id, user_id, title, content, created_at, updated_at
	`

	var post models.Post

	err := DB.QueryRowContext(
		ctx,
		query,
		updatedPost.Title,
		updatedPost.Content,
		id,
		userID,
	).Scan(
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

func DeletePost(ctx context.Context, id int, userID int) error {

	query := `
		DELETE FROM posts 
		WHERE id = $1
		AND user_id = $2
	`

	result, err := DB.ExecContext(ctx, query, id, userID)

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
