package middlewares

import (
	loggerMiddleware "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/log"
)

// Middlewares ...
type Middlewares struct {
	LoggerMiddleware loggerMiddleware.LoggerMiddleware
}

// Bind will bind the service to the middlewares
func (m *Middlewares) Bind(base service.Service) service.Service {
	service := base
	service = m.LoggerMiddleware.Bind(base)
	return service
}

// BindService will bind the service with the middlewares
func BindService(logger log.Logger, service service.Service) service.Service {
	m := Middlewares{
		loggerMiddleware.MakeLoggerMiddleware(logger),
	}
	return m.Bind(service)
}
