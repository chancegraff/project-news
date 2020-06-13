package service

import (
	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/models"
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

// NewService instantiates the service with a connection to the database
func NewService() Service {
	store, err := db.NewStore()
	if err != nil {
		panic(err)
	}
	return &service{store}
}
