package endpoints

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/auth"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/collector"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/token"
	"github.com/chancegraff/project-news/pkg/gateway/service"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	AuthEndpoints      auth.Endpoints
	CollectorEndpoints collector.Endpoints
	RankerEndpoints    ranker.Endpoints
	TokenEndpoints     token.Endpoints
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		AuthEndpoints:      auth.NewEndpoints(s),
		CollectorEndpoints: collector.NewEndpoints(s),
		RankerEndpoints:    ranker.NewEndpoints(s),
		TokenEndpoints:     token.NewEndpoints(s),
	}
}
