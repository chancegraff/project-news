package service

import (
	"time"

	"github.com/chancegraff/project-news/internal/models"
)

// Verify will take an email and password and verify that it matches a record in the database
func (s *service) Verify(email string, password string) (models.User, error) {
	// Get user
	user, err := s.Manager.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// Verify password
	user, err = s.Manager.VerifyPassword(user, password)
	if err != nil {
		return user, err
	}

	// Add verification timestamp
	user, err = s.Manager.AddVerification(user, time.Now())
	if err != nil {
		return user, err
	}

	return user, nil
}
