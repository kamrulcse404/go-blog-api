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

	return  models.Post{}, false
}