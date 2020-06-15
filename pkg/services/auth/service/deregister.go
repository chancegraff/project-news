package service

import (
	"github.com/chancegraff/project-news/internal/models"
)

// Deregister will take a user ID, null its verified_at column, and return that user
func (s *service) Deregister(userID string) (models.User, error) {
	// Search for existing user
	var user models.User
	err := s.Store.Database.First(&user, userID).Error
	if err != nil {
		return user, err
	}

	// Remove verification timestamp
	err = s.Store.Database.Model(&user).Update("verified_at", nil).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
