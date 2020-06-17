package ranker

import (
	pbr "github.com/chancegraff/project-news/api/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
)

// Service implements the collector interface
type Service interface {
	Articles(articleIDs []string) ([]*pbr.ArticleVotes, error)
	User(userID string) (*pbr.UserVotes, error)
	Vote(articleID, userID string) (*pbr.ArticleVotes, error)
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
