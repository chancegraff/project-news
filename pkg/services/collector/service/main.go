package service

import (
	"github.com/chancegraff/project-news/internal/db"
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
	Manager  *manager.Manager
	articles endpoint.Endpoint
}

// NewService instantiates the service with a connection to the database
func NewService() (Service, error) {
	store, err := db.NewStore()
	if err != nil {
		return nil, err
	}
	return &service{
		Manager: &manager.Manager{
			Store: store,
		},
	}, nil
}

// ServiceMiddleware is a chainable middleware for Service
type ServiceMiddleware func(Service) Service
