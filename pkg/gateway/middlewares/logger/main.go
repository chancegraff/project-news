package appLogger

import (
	authLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/auth"
	collectorLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/collector"
	rankerLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/ranker"
	tokenLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/token"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
	"github.com/go-kit/kit/log"
)

// LoggerMiddleware ...
type LoggerMiddleware struct {
	AuthLogger      auth.Middleware
	CollectorLogger collector.Middleware
	RankerLogger    ranker.Middleware
	TokenLogger     token.Middleware
}

// Bind will bind the service to the middlewares
func (m *LoggerMiddleware) Bind(service service.Service) service.Service {
	service.Auth = m.AuthLogger(service.Auth)
	service.Collector = m.CollectorLogger(service.Collector)
	service.Ranker = m.RankerLogger(service.Ranker)
	service.Token = m.TokenLogger(service.Token)
	return service
}

// MakeLoggerMiddleware will create the middlewares
func MakeLoggerMiddleware(lgr log.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		AuthLogger:      authLogger.MakeLoggingMiggleware(lgr),
		CollectorLogger: collectorLogger.MakeLoggingMiggleware(lgr),
		RankerLogger:    rankerLogger.MakeLoggingMiggleware(lgr),
		TokenLogger:     tokenLogger.MakeLoggingMiggleware(lgr),
	}
}
