package middlewares

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/application"
	"github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/transport"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/log"
)

// Middlewares ...
type Middlewares struct {
	ApplicationLogger application.Middleware
	TransportLogger   transport.Middleware
}

// BindEndpoints ...
func (m *Middlewares) BindEndpoints(end endpoints.Endpoints) endpoints.Endpoints {
	endpoints := end
	endpoints = m.TransportLogger.Bind(end)
	return endpoints
}

// BindService ...
func (m *Middlewares) BindService(svc service.Service) service.Service {
	service := svc
	service = m.ApplicationLogger.Bind(svc)
	return service
}

// Bind will bind the service and the endpoints to the middlewares
func (m *Middlewares) Bind(svc service.Service, end endpoints.Endpoints) (service.Service, endpoints.Endpoints) {
	service := m.BindService(svc)
	endpoints := m.BindEndpoints(end)
	return service, endpoints
}

// NewMiddlewares ...
func NewMiddlewares(logger log.Logger) Middlewares {
	return Middlewares{
		application.MakeMiddleware(logger),
		transport.MakeMiddleware(logger),
	}
}
