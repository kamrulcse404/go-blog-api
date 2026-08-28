package handlers

import (
	"blogapi/models"
	"errors"
	"strings"
)

func ValidationPost(post *models.Post) error {

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	if post.Title == "" {
		return errors.New("title is required")
	}

	if len(post.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}

	if len(post.Title) > 255 {
		return errors.New("title cannot exceed 255 characters")
	}

	if post.Content == "" {
		return errors.New("content is required")
	}

	if len(post.Content) < 10 {
		return errors.New("content must be at least 10 characters")
	}

	return nil
}
