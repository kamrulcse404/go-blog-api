package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var postList []Post
var nextID = 1

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API")
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(postList)
}

func post(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	for _, post := range postList {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	post.ID = nextID
	nextID++

	postList = append(postList, post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func main() {
	r := chi.NewRouter()

	r.Get("/users", users)
	r.Get("/posts", posts)
	r.Get("/posts/{id}", post)
	r.Post("/posts", createPost)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
