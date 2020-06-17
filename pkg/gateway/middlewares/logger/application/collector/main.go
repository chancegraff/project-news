package collector

import (
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	next   collector.Service
	logger log.Logger
}

// MakeMiddleware ...
func MakeMiddleware(logger log.Logger) collector.Middleware {
	return func(next collector.Service) collector.Service {
		return &Middleware{
			next,
			log.With(logger, "svc", "collector"),
		}
	}
}
