package token

import (
	pbt "github.com/chancegraff/project-news/api/token"
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
)

// Service implements the collector interface
type Service interface {
	Generate(identifiers *pbt.Identifiers, client *pbt.Client) (string, error)
	Verify(identifiers *pbt.Identifiers, client *pbt.Client) (string, error)
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
