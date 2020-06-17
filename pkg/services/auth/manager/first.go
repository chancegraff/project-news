package manager

import (
	"github.com/chancegraff/project-news/internal/models"
)

// FindByID will search for and return a user with the ID provided
func (m *Manager) FindByID(userID string) (models.User, error) {
	var user models.User
	err := m.Store.Database.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindByEmail will search for and return a user with the ID provided
func (m *Manager) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := m.Store.Database.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
