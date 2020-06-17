package manager

import (
	"time"

	"github.com/chancegraff/project-news/internal/models"
)

// AddVerification will take a user and a timestamp and update that users verified_at with thit
func (m *Manager) AddVerification(user models.User, timestamp time.Time) (models.User, error) {
	err := m.Store.Database.Model(&user).Update("verified_at", timestamp).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// RemoveVerification will take a user and update the users verified_at with nil
func (m *Manager) RemoveVerification(user models.User) (models.User, error) {
	err := m.Store.Database.Model(&user).Update("verified_at", nil).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
