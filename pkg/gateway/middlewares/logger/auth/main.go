package auth

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	next   auth.Service
	logger log.Logger
}

// MakeLoggingMiggleware ...
func MakeLoggingMiggleware(logger log.Logger) auth.Middleware {
	return func(next auth.Service) auth.Service {
		return &LoggingMiddleware{
			next,
			logger,
		}
	}
}
