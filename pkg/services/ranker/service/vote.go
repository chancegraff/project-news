package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/jinzhu/gorm"
)

// Vote will take an article ID and a user ID and returns the new vote count
func (s *service) Vote(articleID, userID string) (models.ArticleVotes, error) {
	articleVotes := models.ArticleVotes{ArticleID: articleID}
	buffer := models.Vote{
		ArticleID: articleID,
		UserID:    userID,
	}

	// Look for an existing record and check for errors
	findErr := s.Store.Database.Where(models.Vote{UserID: userID, ArticleID: articleID}).First(&buffer).Error
	if gorm.IsRecordNotFoundError(findErr) {
		// If the record does not exist, create it
		err := s.Store.Database.Create(&buffer).Error
		if err != nil {
			return articleVotes, err
		}
	} else {
		// If the record does exist, delete it
		err := s.Store.Database.Delete(&buffer).Error
		if err != nil {
			return articleVotes, err
		}
	}

	// Return array of user IDs associated with article
	err := s.Store.Database.Select("article_id, count(*) as votes").Where("article_id = ?", articleID).Find(&articleVotes).Error
	if err != nil {
		return articleVotes, err
	}

	return articleVotes, nil
}
