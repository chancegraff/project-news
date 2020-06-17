package token

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	next   token.Service
	logger log.Logger
}

// MakeLoggingMiggleware ...
func MakeLoggingMiggleware(logger log.Logger) token.Middleware {
	return func(next token.Service) token.Service {
		return &LoggingMiddleware{
			next,
			logger,
		}
	}
}
