package main

import (
	"fmt"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API")
}

func posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Posts API")
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", users)
	mux.HandleFunc("/posts", posts)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
