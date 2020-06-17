package token

import (
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	GenerateEndpoint endpoint.Endpoint
	VerifyEndpoint   endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GenerateEndpoint: MakeGenerateEndpoint(s),
		VerifyEndpoint:   MakeVerifyEndpoint(s),
	}
}

// Middleware is a chainable middleware for Endpoints
type Middleware func(Endpoints) Endpoints
