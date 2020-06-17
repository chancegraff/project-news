package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/auth/manager"
)

// Service implements the collector interface
type Service interface {
	Deregister(userID string) (models.User, error)
	Register(email string, password string) (models.User, error)
	User(userID string) (models.User, error)
	Verify(email string, password string) (models.User, error)
}

type service struct {
	Manager *manager.Manager
}

// NewService instantiates the service with a connection to the database
func NewService(mgr *manager.Manager) Service {
	return &service{
		Manager: mgr,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
