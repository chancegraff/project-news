package service

import "github.com/chancegraff/project-news/internal/models"

// User will take a user ID and returns an array of article IDs associated with it
func (s *service) User(userID string) (models.UserVotes, error) {
	userVotes := models.UserVotes{UserID: userID}

	// Get votes
	err := s.Store.Database.Select("*").Where("user_id = ?", userID).Find(&userVotes.Votes).Error
	if err != nil {
		return userVotes, err
	}

	return userVotes, nil
}
