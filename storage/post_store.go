package storage

import "blogapi/models"

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

	var posts []models.Post
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

// func UpdatePost(id int, updatedPost models.Post) (models.Post, bool) {

// 	for i, post := range postList {
// 		if post.ID == id {
// 			postList[i].Title = updatedPost.Title
// 			postList[i].Content = updatedPost.Content
// 			return postList[i], true
// 		}
// 	}

// 	return models.Post{}, false
// }

// func DeletePost(id int) bool {
// 	for i, post := range postList {
// 		if post.ID == id {
// 			postList = append(postList[:i], postList[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }
