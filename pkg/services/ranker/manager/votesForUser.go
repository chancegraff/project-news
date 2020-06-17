package manager

import "github.com/chancegraff/project-news/internal/models"

// GetVotesForUser will take a user ID and return the articles that user has voted on
func (m *Manager) GetVotesForUser(userID string) (models.UserVotes, error) {
	userVotes := models.UserVotes{UserID: userID}

	// Get votes
	err := m.Store.Database.Select("*").Where("user_id = ?", userID).Find(&userVotes.Votes).Error
	if err != nil {
		return userVotes, err
	}

	return userVotes, nil
}
