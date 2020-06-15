package service

import (
	"github.com/chancegraff/project-news/internal/models"
)

// Deregister will take a user ID, null its verified_at column, and return that user
func (s *service) Deregister(userID string) (models.User, error) {
	// Search for existing user
	user, err := s.Manager.FindByID(userID)
	if err != nil {
		return user, err
	}

	// Remove verification timestamp
	user, err = s.Manager.RemoveVerification(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
