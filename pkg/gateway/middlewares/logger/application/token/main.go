package token

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	next   token.Service
	logger log.Logger
}

// MakeMiddleware ...
func MakeMiddleware(logger log.Logger) token.Middleware {
	return func(next token.Service) token.Service {
		return &Middleware{
			next,
			log.With(logger, "svc", "token"),
		}
	}
}
