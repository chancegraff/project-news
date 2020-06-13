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
	manager  *manager.Manager
	articles endpoint.Endpoint
}

// NewService instantiates the service with a connection to the database
func NewService() Service {
	store, err := db.NewStore()
	if err != nil {
		panic(err)
	}
	svc := &service{
		manager: &manager.Manager{
			Store: store,
		},
	}
	go svc.Collect()
	return svc
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
