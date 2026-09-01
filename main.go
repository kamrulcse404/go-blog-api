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

	defer storage.DB.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)

	r.Get("/posts", handlers.Posts)
	r.Get("/posts/{id}", handlers.Post)

	r.Post("/users/register", handlers.RegisterUser)
	r.Post("/users/login", handlers.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.Get("/users/me", handlers.GetCurrentUser)

		r.Post("/posts", handlers.CreatePost)
		r.Put("/posts/{id}", handlers.UpdatePost)
		r.Delete("/posts/{id}", handlers.DeletePost)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server running on : 8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server Failed: ", err)
	}
}
