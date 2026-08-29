package handlers

import (
	"blogapi/models"
	"blogapi/security"
	"blogapi/storage"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var req models.RegisterUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = ValidateRegisterRequest(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := security.HashPassword(req.Password)

	if err != nil {
		log.Printf("failed to hash password: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	user, err = storage.CreateUser(r.Context(), user)
	if err != nil {

		if errors.Is(err, storage.ErrEmailAlreadyExists) {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}

		log.Printf("failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
