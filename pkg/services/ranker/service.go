package ranker

import (
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/models"
	"github.com/jinzhu/gorm"
)

// Service ...
type Service interface {
	Articles(articleIDs []string) ([]models.ArticleVotes, error)
	User(userID string) (models.UserVotes, error)
	Vote(articleID, userID string) (models.ArticleVotes, error)
}

type service struct {
	Store *db.Store
}

// Articles will take an array of article IDs and return the votes for each
func (s *service) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	var articleVotesArray []models.ArticleVotes
	yesterday := time.Now().AddDate(0, 0, -1)
	db := s.Store.Database

	// Build query
	subquery := db.
		Select("vote.*").
		Table("vote").
		SubQuery()
	query := db.
		Select("vote.article_id, count(t1.*) as votes").
		Joins("LEFT JOIN ? AS t1 ON t1.article_id = vote.article_id", subquery).
		Where("vote.article_id IN (?)", articleIDs).
		Where("vote.created_at > ?", yesterday).
		Group("vote.article_id")

	// Commit transaction
	if err := query.Scan(&articleVotesArray).Error; err != nil {
		return articleVotesArray, err
	}

	return articleVotesArray, nil
}

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

// NewService instantiates the service with a connection to the database
func newService() (*service, error) {
	store, err := db.NewStore()
	if err != nil {
		return nil, err
	}
	return &service{store}, nil
}
