package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/auth"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/collector"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/token"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// Middleware ...
type Middleware struct {
	AuthLogger      auth.Middleware
	CollectorLogger collector.Middleware
	RankerLogger    ranker.Middleware
	TokenLogger     token.Middleware
}

// Bind will bind the endpoints to the middlewares
func (m *Middleware) Bind(endpoints endpoints.Endpoints) endpoints.Endpoints {
	endpoints.AuthEndpoints = m.AuthLogger(endpoints.AuthEndpoints)
	endpoints.CollectorEndpoints = m.CollectorLogger(endpoints.CollectorEndpoints)
	endpoints.RankerEndpoints = m.RankerLogger(endpoints.RankerEndpoints)
	endpoints.TokenEndpoints = m.TokenLogger(endpoints.TokenEndpoints)
	return endpoints
}

// MakeMiddleware will create the middlewares
func MakeMiddleware(logger log.Logger) Middleware {
	lgr := log.With(logger, "logger", "transport")
	return Middleware{
		AuthLogger:      MakeAuthMiddleware(lgr),
		CollectorLogger: MakeCollectorMiddleware(lgr),
		RankerLogger:    MakeRankerMiddleware(lgr),
		TokenLogger:     MakeTokenMiddleware(lgr),
	}
}

// MakeEndpoint ...
func MakeEndpoint(method string, logger log.Logger, next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (output interface{}, err error) {
		logger.Log(
			"start", "start",
			"method", method,
			"request", fmt.Sprint(request),
			"output", fmt.Sprint(output),
			"err", err,
		)
		defer func(begin time.Time) {
			logger.Log(
				"method", method,
				"request", fmt.Sprint(request),
				"output", fmt.Sprint(output),
				"err", err,
				"took", time.Since(begin),
			)
		}(time.Now())
		output, err = next(ctx, request)
		return
	}
}
