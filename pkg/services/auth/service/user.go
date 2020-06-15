package service

import "github.com/chancegraff/project-news/internal/models"

// User will take a user ID and return that record from the database
func (s *service) User(userID string) (models.User, error) {
	// Get user
	var user models.User
	err := s.Store.Database.Select("article_id").Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
