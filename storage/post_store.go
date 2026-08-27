package storage

import (
	"blogapi/models"
	"database/sql"
)

func GetPosts() ([]models.Post, error) {

	rows, err := DB.Query(`
		SELECT id, title, content
		FROM posts
		ORDER BY id
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post

		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
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

func CreatePost(post models.Post) (models.Post, error) {
	query := `
		INSERT INTO posts (title, content)
		VALUES ($1, $2)
		RETURNING id
	`

	err := DB.QueryRow(query, post.Title, post.Content).Scan(&post.ID)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetPostByID(id int) (models.Post, error) {
	query := `
		SELECT id, title, content 
		FROM posts 
		WHERE id = $1
	`
	var post models.Post

	err := DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
	)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func UpdatePost(id int, updatedPost models.Post) (models.Post, error) {

	query := `
		UPDATE posts 
		set title = $1, content = $2
		WHERE id = $3
		RETURNING id, title, content
	`
	var post models.Post

	err := DB.QueryRow(query, updatedPost.Title, updatedPost.Content, id).Scan(&post.ID, &post.Title, &post.Content)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func DeletePost(id int) error {
	
	query := `
		DELETE FROM posts 
		WHERE id = $1
	`

	result, err := DB.Exec(query, id)

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
