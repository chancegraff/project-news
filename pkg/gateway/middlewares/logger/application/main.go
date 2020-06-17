package application

import (
	authLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/application/auth"
	collectorLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/application/collector"
	rankerLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/application/ranker"
	tokenLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/application/token"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	AuthLogger      auth.Middleware
	CollectorLogger collector.Middleware
	RankerLogger    ranker.Middleware
	TokenLogger     token.Middleware
}

// Bind will bind the service to the middlewares
func (m *Middleware) Bind(service service.Service) service.Service {
	service.Auth = m.AuthLogger(service.Auth)
	service.Collector = m.CollectorLogger(service.Collector)
	service.Ranker = m.RankerLogger(service.Ranker)
	service.Token = m.TokenLogger(service.Token)
	return service
}

// MakeMiddleware will create the middlewares
func MakeMiddleware(logger log.Logger) Middleware {
	lgr := log.With(logger, "logger", "application")
	return Middleware{
		AuthLogger:      authLogger.MakeMiddleware(lgr),
		CollectorLogger: collectorLogger.MakeMiddleware(lgr),
		RankerLogger:    rankerLogger.MakeMiddleware(lgr),
		TokenLogger:     tokenLogger.MakeMiddleware(lgr),
	}
}
