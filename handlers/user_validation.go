package handlers

import (
	"blogapi/models"
	"errors"
	"net/mail"
	"strings"
)

func ValidateRegisterRequest(req *models.RegisterUserRequest) error {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Name == "" {
		return errors.New("name is required")
	}

	if len(req.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if len(req.Name) > 100 {
		return errors.New("name cannot exceed 100 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	if len(req.Email) > 255 {
		return errors.New("email cannot exceed 255 characters")
	}

	parsedEmail, err := mail.ParseAddress(req.Email)
	if err != nil || parsedEmail.Address != req.Email {
		return errors.New("invalid email address")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(req.Password) > 72 {
		return errors.New("password cannot exceed 72 bytes")
	}

	return nil
}

func ValidateLoginRequest(req *models.LoginRequest) error {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Email == "" {
		return errors.New("email is required")
	}

	if len(req.Email) > 255 {
		return errors.New("email cannot exceed 255 characters")
	}

	parsedEmail, err := mail.ParseAddress(req.Email)
	if err != nil || parsedEmail.Address != req.Email {
		return errors.New("invalid email address")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(req.Password) > 72 {
		return errors.New("password cannot exceed 72 bytes")
	}

	return nil
}
