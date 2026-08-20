package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Blog API is running")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on : 8080")

	server.ListenAndServe()
}
