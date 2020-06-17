package auth

import (
	pba "github.com/chancegraff/project-news/api/auth"
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
)

// Service implements the auth interface
type Service interface {
	Deregister(userID string) (*pba.User, error)
	Register(email string, password string) (*pba.User, error)
	User(userID string) (*pba.User, error)
	Verify(email string, password string) (*pba.User, error)
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
