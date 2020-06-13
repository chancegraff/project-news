package middlewares

import "github.com/chancegraff/project-news/pkg/services/collector/service"

// Middlewares ...
type Middlewares struct {
	Articles service.Middleware
}

// Bind will bind the service to the middlewares
func (m *Middlewares) Bind(base service.Service) service.Service {
	service := base
	service = m.Articles(base)
	return service
}

// BindService will bind the service with the middlewares
func BindService(service service.Service) service.Service {
	m := Middlewares{
		ArticlesProxyMiddleware(),
	}
	return m.Bind(service)
}
