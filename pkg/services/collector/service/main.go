package service

import (
	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/collector/manager"
	"github.com/go-kit/kit/endpoint"
)

// Service implements the collector interface
type Service interface {
	All(offset int) ([]models.Article, error)
	Get(id int) (models.Article, error)
}

type service struct {
	manager          *manager.Manager
	ArticlesEndpoint endpoint.Endpoint
}

// NewService instantiates the service with a connection to the database
func NewService(manager *manager.Manager) Service {
	return &service{manager: manager}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
