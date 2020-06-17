package manager

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/jinzhu/gorm"
)

// Get will look for an existing vote and return it
func (m *Manager) Get(articleID, userID string) (models.Vote, error) {
	vote := models.Vote{}
	err := m.Store.Database.Where("user_id = ?", userID).Where("article_id = ?", articleID).First(&vote).Error
	if err != nil {
		return vote, err
	}
	return vote, nil
}

// CreateOrDelete will look for an existing vote and delete it if it exists or create a new one if not
func (m *Manager) CreateOrDelete(articleID, userID string) error {
	vote, err := m.Get(articleID, userID)
	if gorm.IsRecordNotFoundError(err) {
		return m.Create(vote)
	}
	return m.Delete(vote)
}

// Create will save a new instance of vote to the database
func (m *Manager) Create(vote models.Vote) error {
	return m.Store.Database.Create(&vote).Error
}

// Delete will remove an existing instance of vote from the database
func (m *Manager) Delete(vote models.Vote) error {
	return m.Store.Database.Delete(&vote).Error
}
