package handlers

import (
	"fmt"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users API")
}
