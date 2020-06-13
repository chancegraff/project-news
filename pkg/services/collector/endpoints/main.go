package endpoints

import (
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	All endpoint.Endpoint
	Get endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		All: makeAllEndpoint(s),
		Get: makeGetEndpoint(s),
	}
}
