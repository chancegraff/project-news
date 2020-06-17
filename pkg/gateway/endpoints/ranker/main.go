package ranker

import (
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	ArticlesEndpoint endpoint.Endpoint
	UserEndpoint     endpoint.Endpoint
	VoteEndpoint     endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		ArticlesEndpoint: MakeArticlesEndpoint(s),
		UserEndpoint:     MakeUserEndpoint(s),
		VoteEndpoint:     MakeVoteEndpoint(s),
	}
}

// Middleware is a chainable middleware for Endpoints
type Middleware func(Endpoints) Endpoints
