package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Register will take credentials, hash them, store them in the database, and return the new record
func (s *service) Register(email string, password string) (models.User, error) {
	// Check user does not exist
	var existing models.User
	err := s.Store.Database.Where("email = ?", email).First(&existing).Error
	if err != nil {
		return existing, err
	}

	// Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	// Create new user
	user := models.User{
		Credentials: models.Credentials{
			Email:    email,
			Password: string(hash),
		},
	}

	// Create user
	err = s.Store.Database.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
