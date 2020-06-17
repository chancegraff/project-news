package transport

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/collector"
	"github.com/go-kit/kit/log"
)

// MakeCollectorMiddleware ...
func MakeCollectorMiddleware(logger log.Logger) collector.Middleware {
	lgr := log.With(logger, "end", "collector")
	return func(next collector.Endpoints) collector.Endpoints {
		return collector.Endpoints{
			AllEndpoint: MakeEndpoint("All", lgr, next.AllEndpoint),
			GetEndpoint: MakeEndpoint("Get", lgr, next.GetEndpoint),
		}
	}
}
