package collector

import (
	pbc "github.com/chancegraff/project-news/api/collector"
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
)

// Service implements the collector interface
type Service interface {
	All(offset int) ([]*pbc.Article, error)
	Get(id int) (*pbc.Article, error)
}

type service struct {
	Proxy *proxy.Proxy
}

// NewService instantiates the service
func NewService(prx *proxy.Proxy) Service {
	return &service{
		Proxy: prx,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
