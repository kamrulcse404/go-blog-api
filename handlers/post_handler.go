package handlers

import (
	"blogapi/middleware"
	"blogapi/models"
	"blogapi/storage"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Posts(w http.ResponseWriter, r *http.Request) {

	limit := 10
	offset := 0

	limitValue := r.URL.Query().Get("limit")
	search := strings.TrimSpace(r.URL.Query().Get("search"))

	userID := 0

	userIDValue := r.URL.Query().Get("user_id")

	// filter By User ID
	if userIDValue != "" {
		parsed, err := strconv.Atoi(userIDValue)
		if err != nil || parsed <= 0 {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}
		userID = parsed
	}

	// limit 
	if limitValue != "" {
		parsed, err := strconv.Atoi(limitValue)
		if err != nil || parsed <= 0 {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}

		if parsed > 100 {
			http.Error(w, "Limit cannot exceed 100", http.StatusBadRequest)
			return
		}

		limit = parsed
	}

	// offset 
	if value := r.URL.Query().Get("offset"); value != "" {
		parsed, err := strconv.Atoi(value)

		if err != nil || parsed < 0 {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}

		offset = parsed
	}

	ctx := r.Context()
	posts, err := storage.GetPosts(ctx, limit, offset, search, userID)

	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	total, err := storage.CountPosts(ctx, search, userID)
	if err != nil {
		http.Error(w, "Failed to count posts", http.StatusInternalServerError)
		return
	}

	response := models.PostListResponse{
		Data: posts,
		Pagination: models.Pagination{
			Limit:  limit,
			Offset: offset,
			Total:  total,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = ValidatePost(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.UserID = userID
	post, err = storage.CreatePost(r.Context(), post)

	if err != nil {
		log.Printf("failed to create post: %v", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func Post(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := storage.GetPostByID(r.Context(), id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

	userID, ok := middleware.GetUserID(r.Context())

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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

	err = ValidatePost(&updatedPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := storage.UpdatePost(r.Context(), id, userID, updatedPost)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

func DeletePost(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = storage.DeletePost(r.Context(), id, userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		log.Printf("failed to delete post %d: %v", id, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
