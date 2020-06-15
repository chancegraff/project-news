package manager

import (
	"github.com/chancegraff/project-news/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword will take a user instance and a password string and return an error if the password is incorrect
func (m *Manager) VerifyPassword(user models.User, password string) (models.User, error) {
	// Verify passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
