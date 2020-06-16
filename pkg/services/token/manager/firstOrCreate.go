package manager

import (
	"time"

	"github.com/chancegraff/project-news/internal/models"
)

// FirstOrCreate ...
func (m *Manager) FirstOrCreate(client models.Client) (models.Client, error) {
	buffer := client
	err := m.Store.Database.Where("hash = ? AND expired_at > ?", client.Hash, time.Now()).FirstOrCreate(&buffer).Error
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
