package manager

import (
	"time"

	"github.com/chancegraff/project-news/internal/models"
)

// First ...
func (m *Manager) First(client models.Client) (models.Client, error) {
	buffer := client
	err := m.Store.Database.Where("hash = ? AND expired_at > ?", client.Hash, time.Now()).First(&buffer).Error
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
