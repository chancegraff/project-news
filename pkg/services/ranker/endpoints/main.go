package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
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
