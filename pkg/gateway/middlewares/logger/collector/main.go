package collector

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	next   collector.Service
	logger log.Logger
}

// MakeLoggingMiggleware ...
func MakeLoggingMiggleware(logger log.Logger) collector.Middleware {
	return func(next collector.Service) collector.Service {
		return &LoggingMiddleware{
			next,
			logger,
		}
	}
}
