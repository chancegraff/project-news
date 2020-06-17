package middlewares

import (
	"os"

	authLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/auth"
	collectorLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/collector"
	rankerLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/ranker"
	tokenLogger "github.com/chancegraff/project-news/pkg/gateway/middlewares/logger/token"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/chancegraff/project-news/pkg/gateway/service/auth"
	"github.com/chancegraff/project-news/pkg/gateway/service/collector"
	"github.com/chancegraff/project-news/pkg/gateway/service/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/service/token"
	"github.com/go-kit/kit/log"
)

// Middlewares ...
type Middlewares struct {
	AuthLogger      auth.Middleware
	CollectorLogger collector.Middleware
	RankerLogger    ranker.Middleware
	TokenLogger     token.Middleware
}

// Bind will bind the service to the middlewares
func (m *Middlewares) Bind(service service.Service) service.Service {
	service.Auth = m.AuthLogger(service.Auth)
	service.Collector = m.CollectorLogger(service.Collector)
	service.Ranker = m.RankerLogger(service.Ranker)
	service.Token = m.TokenLogger(service.Token)
	return service
}

// NewMiddleware will create the middlewares
func NewMiddleware(service service.Service) (Middlewares, log.Logger) {
	lgr := log.NewLogfmtLogger(os.Stderr)
	lgr = log.With(lgr, "ts", log.DefaultTimestampUTC)
	lgr = log.With(lgr, "caller", log.DefaultCaller)
	lgr = log.With(lgr, "service", "gateway")
	return Middlewares{
		AuthLogger:      authLogger.MakeLoggingMiggleware(lgr),
		CollectorLogger: collectorLogger.MakeLoggingMiggleware(lgr),
		RankerLogger:    rankerLogger.MakeLoggingMiggleware(lgr),
		TokenLogger:     tokenLogger.MakeLoggingMiggleware(lgr),
	}, lgr
}
