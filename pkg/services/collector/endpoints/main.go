package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	AllEndpoint endpoint.Endpoint
	GetEndpoint endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		AllEndpoint: MakeAllEndpoint(s),
		GetEndpoint: MakeGetEndpoint(s),
	}
}
