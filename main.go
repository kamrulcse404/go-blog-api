package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API")
}

func posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Posts API")
}


func main() {
	r := chi.NewRouter()

	r.Get("/users", users)
	r.Get("/posts", posts)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
