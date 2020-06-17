package ranker

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	next   ranker.Service
	logger log.Logger
}

// MakeMiddleware ...
func MakeMiddleware(logger log.Logger) ranker.Middleware {
	return func(next ranker.Service) ranker.Service {
		return &Middleware{
			next,
			log.With(logger, "svc", "ranker"),
		}
	}
}
