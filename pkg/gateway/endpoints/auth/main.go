package auth

import (
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	DeregisterEndpoint endpoint.Endpoint
	RegisterEndpoint   endpoint.Endpoint
	UserEndpoint       endpoint.Endpoint
	VerifyEndpoint     endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		DeregisterEndpoint: MakeDeregisterEndpoint(s),
		RegisterEndpoint:   MakeRegisterEndpoint(s),
		UserEndpoint:       MakeUserEndpoint(s),
		VerifyEndpoint:     MakeVerifyEndpoint(s),
	}
}
