package main

import (
	"blogapi/handlers"
	"blogapi/middleware"
	"blogapi/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	if err := storage.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

	// r.Get("/users", handlers.Users)
	r.Get("/posts", handlers.Posts)
	r.Post("/posts", handlers.CreatePost)
	// r.Get("/posts/{id}", handlers.Post)
	// r.Put("/posts/{id}", handlers.UpdatePost)
	// r.Delete("/posts/{id}", handlers.DeletePost)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
