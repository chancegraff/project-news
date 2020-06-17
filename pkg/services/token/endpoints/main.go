package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/token/service"
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
