package service

import (
	"github.com/chancegraff/project-news/internal/models"
)

// Vote will take an article ID and a user ID and returns the new vote count
func (s *service) Vote(articleID, userID string) (models.ArticleVotes, error) {
	// Create a new vote or delete an existing one
	err := s.Manager.CreateOrDelete(articleID, userID)
	if err != nil {
		return models.ArticleVotes{}, err
	}

	// Get the new vote count for the article
	articleVotes, err := s.Manager.GetVotesForArticle(articleID)
	if err != nil {
		return models.ArticleVotes{}, err
	}

	return articleVotes, nil
}
