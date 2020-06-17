package service

import "github.com/chancegraff/project-news/internal/models"

// User will take a user ID and return that record from the database
func (s *service) User(userID string) (models.User, error) {
	return s.Manager.FindByID(userID)
}
