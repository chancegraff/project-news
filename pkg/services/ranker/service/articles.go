package service

import (
	"log"

	"github.com/chancegraff/project-news/internal/models"
)

// Articles will take an array of article IDs and return the votes for each
func (s *service) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	var articleVotesArray []models.ArticleVotes
	db := s.Manager.Store.Database

	// Build query
	query := db.
		Table("votes").
		Select("article_id, count(*) as votes").
		Where("article_id IN (?)", articleIDs).
		Group("article_id")

	// Commit transaction
	log.Println("Querying for articles", query.QueryExpr())
	if err := query.Find(&articleVotesArray).Error; err != nil {
		return articleVotesArray, err
	}

	return articleVotesArray, nil
}
