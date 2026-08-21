package storage

import "blogapi/models"

var postList []models.Post
var nextID = 1

func GetPosts() []models.Post {

	return postList
}

func CreatePost(post models.Post) models.Post {
	post.ID = nextID
	nextID++

	postList = append(postList, post)
	return post
}

func GetPostByID(id int) (models.Post, bool) {

	for _, post := range postList {
		if post.ID == id {
			return post, true
		}
	}

	return models.Post{}, false
}

func UpdatePost(id int, updatedPost models.Post) (models.Post, bool) {

	for i, post := range postList {
		if post.ID == id {
			postList[i].Title = updatedPost.Title
			postList[i].Content = updatedPost.Content
			return postList[i], true
		}
	}

	return models.Post{}, false
}

func DeletePost(id int) bool {
	for i, post := range postList {
		if post.ID == id {
			postList = append(postList[:i], postList[i+1:]...)
			return true
		}
	}
	return false
}
