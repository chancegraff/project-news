package service

import (
	"context"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/models"
)

// Service implements the collector interface
type Service interface {
	Deregister(userID string) (models.User, error)
	Register(email string, password string) (models.User, error)
	User(userID string) (models.User, error)
	Verify(email string, password string) (models.User, error)
}

type service struct {
	Store *db.Store
}

// NewService instantiates the service with a connection to the database
func NewService(ctx context.Context, store *db.Store) Service {
	return &service{
		Store: store,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
