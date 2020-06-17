package ranker

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	next   ranker.Service
	logger log.Logger
}

// MakeLoggingMiggleware ...
func MakeLoggingMiggleware(logger log.Logger) ranker.Middleware {
	return func(next ranker.Service) ranker.Service {
		return &LoggingMiddleware{
			next,
			logger,
		}
	}
}
