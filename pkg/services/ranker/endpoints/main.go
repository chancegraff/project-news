package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	Articles endpoint.Endpoint
	User     endpoint.Endpoint
	Vote     endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Articles: makeArticlesEndpoint(s),
		User:     makeUserEndpoint(s),
		Vote:     makeVoteEndpoint(s),
	}
}
