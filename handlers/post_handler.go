package handlers

import (
	"blogapi/models"
	"blogapi/storage"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	posts, err := storage.GetPosts()

	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	if post.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if post.Content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	post, err = storage.CreatePost(post)

	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func Post(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := storage.GetPostByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updatedPost models.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updatedPost.Title = strings.TrimSpace(updatedPost.Title)
	updatedPost.Content = strings.TrimSpace(updatedPost.Content)

	if updatedPost.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if updatedPost.Content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	post, err := storage.UpdatePost(id, updatedPost)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		log.Printf("failed to update post %d: %v", id, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)

	if err != nil {
		log.Printf("failed to encode post response: %v", err)
	}
}

// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	idParam := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		http.Error(w, "Invalid post ID", http.StatusBadRequest)
// 		return
// 	}

// 	deleted := storage.DeletePost(id)

// 	if !deleted {
// 		http.Error(w, "Post not found", http.StatusNotFound)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
