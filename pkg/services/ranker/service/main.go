package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/ranker/manager"
)

// Service implements the ranker interface
type Service interface {
	Articles(articleIDs []string) ([]models.ArticleVotes, error)
	User(userID string) (models.UserVotes, error)
	Vote(articleID, userID string) (models.ArticleVotes, error)
}

type service struct {
	Manager *manager.Manager
}

// NewService instantiates the service with a connection to the database
func NewService(mgr *manager.Manager) Service {
	return &service{Manager: mgr}
}
