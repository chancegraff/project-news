package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Verify will take an email and password and verify that it matches a record in the database
func (s *service) Verify(email string, password string) (models.User, error) {
	// Get user from database
	var user models.User
	err := s.Store.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	// Verify passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
