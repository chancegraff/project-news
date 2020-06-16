package middlewares

import (
	"github.com/chancegraff/project-news/pkg/services/token/service"
	"github.com/go-kit/kit/log"
)

// Middlewares ...
type Middlewares struct {
	Logger service.Middleware
}

// Bind will bind the service to the middlewares
func (m *Middlewares) Bind(base service.Service) service.Service {
	service := base
	service = m.Logger(base)
	return service
}

// BindService will bind the service with the middlewares
func BindService(logger log.Logger, service service.Service) service.Service {
	m := Middlewares{
		MakeLoggingMiggleware(logger),
	}
	return m.Bind(service)
}
