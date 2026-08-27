package storage

import "blogapi/models"

// var postList []models.Post
// var nextID = 1

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

// func CreatePost(post models.Post) models.Post {
// 	post.ID = nextID
// 	nextID++

// 	postList = append(postList, post)
// 	return post
// }

// func GetPostByID(id int) (models.Post, bool) {

// 	for _, post := range postList {
// 		if post.ID == id {
// 			return post, true
// 		}
// 	}

// 	return models.Post{}, false
// }

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
