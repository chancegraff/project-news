package auth

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	next   auth.Service
	logger log.Logger
}

// MakeMiddleware ...
func MakeMiddleware(logger log.Logger) auth.Middleware {
	return func(next auth.Service) auth.Service {
		return &Middleware{
			next,
			log.With(logger, "svc", "auth"),
		}
	}
}
