package service

import (
	"time"

	"github.com/chancegraff/project-news/internal/models"
)

// Articles will take an array of article IDs and return the votes for each
func (s *service) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	var articleVotesArray []models.ArticleVotes
	yesterday := time.Now().AddDate(0, 0, -1)
	db := s.Store.Database

	// Build query
	query := db.
		Select("distinct(v1.article_id), count(v2.user_id) as votes").
		Joins("LEFT JOIN vote AS v2 ON v2.article_id = v1.article_id").
		Where("v1.article_id IN (?)", articleIDs).
		Where("v1.created_at > ?", yesterday).
		Order("v1 desc").
		Group("votes")

	// Commit transaction
	if err := query.Scan(&articleVotesArray).Error; err != nil {
		return articleVotesArray, err
	}

	return articleVotesArray, nil
}
