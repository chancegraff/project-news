package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	DeregisterEndpoint endpoint.Endpoint
	RegisterEndpoint   endpoint.Endpoint
	UserEndpoint       endpoint.Endpoint
	VerifyEndpoint     endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		DeregisterEndpoint: MakeDeregisterEndpoint(s),
		RegisterEndpoint:   MakeRegisterEndpoint(s),
		UserEndpoint:       MakeUserEndpoint(s),
		VerifyEndpoint:     MakeVerifyEndpoint(s),
	}
}
