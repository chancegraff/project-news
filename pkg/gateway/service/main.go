package service

import (
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
)

// Service implements the collector interface
type Service struct {
	Auth      auth.Service
	Collector collector.Service
	Ranker    ranker.Service
	Token     token.Service
}

// NewService instantiates the service with a connection to the database
func NewService(prx *proxy.Proxy) Service {
	return Service{
		Auth:      auth.NewService(prx),
		Collector: collector.NewService(prx),
		Ranker:    ranker.NewService(prx),
		Token:     token.NewService(prx),
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
