package service

import (
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Register will take credentials, hash them, store them in the database, and return the new record
func (s *service) Register(email string, password string) (models.User, error) {
	// Check email does not exist
	existing, err := s.Manager.FindByEmail(email)
	if !gorm.IsRecordNotFoundError(err) {
		return existing, errors.New("Record already exists")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return models.User{}, err
	}

	// Create user
	user, err := s.Manager.Create(email, hash)
	if err != nil {
		return user, err
	}

	return user, nil
}
