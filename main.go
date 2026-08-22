package main

import (
	"blogapi/handlers"
	"blogapi/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

	r.Get("/users", handlers.Users)
	r.Get("/posts", handlers.Posts)
	r.Get("/posts/{id}", handlers.Post)
	r.Post("/posts", handlers.CreatePost)
	r.Put("/posts/{id}", handlers.UpdatePost)
	r.Delete("/posts/{id}", handlers.DeletePost)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
