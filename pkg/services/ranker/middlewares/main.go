package middlewares

import "github.com/chancegraff/project-news/pkg/services/ranker/service"

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
func BindService(service service.Service) service.Service {
	m := Middlewares{
		MakeLoggingMiggleware(),
	}
	return m.Bind(service)
}
