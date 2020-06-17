package service

import (
	"github.com/chancegraff/project-news/internal/models"
)

// Articles will take an array of article IDs and return the votes for each
func (s *service) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	return s.Manager.GetVotesForArticles(articleIDs)
}
