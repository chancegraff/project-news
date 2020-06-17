package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/token/manager"
)

// Service implements the ranker interface
type Service interface {
	Generate(identifiers models.Identifiers, client models.Client) (string, error)
	Verify(identifiers models.Identifiers, client models.Client) (string, error)
}

type service struct {
	Manager *manager.Manager
}

// NewService instantiates the service with a connection to the database
func NewService(mgr *manager.Manager) Service {
	return &service{Manager: mgr}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
