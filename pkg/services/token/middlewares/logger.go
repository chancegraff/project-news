package middlewares

import (
	"fmt"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/token/service"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	next   service.Service
	logger log.Logger
}

// MakeLoggingMiggleware ...
func MakeLoggingMiggleware(logger log.Logger) service.Middleware {
	return func(next service.Service) service.Service {
		return &LoggingMiddleware{
			next,
			logger,
		}
	}
}

// Generate ...
func (mw *LoggingMiddleware) Generate(identifiers models.Identifiers, client models.Client) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"identifiers", fmt.Sprint(identifiers),
			"client", fmt.Sprint(client),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Generate(identifiers, client)
	return
}

// Verify ...
func (mw *LoggingMiddleware) Verify(identifiers models.Identifiers, client models.Client) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"identifiers", fmt.Sprint(identifiers),
			"client", fmt.Sprint(client),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Verify(identifiers, client)
	return
}
