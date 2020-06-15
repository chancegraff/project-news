package middlewares

import (
	"fmt"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
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

// All ...
func (mw *LoggingMiddleware) All(offset int) (output []models.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"offset", fmt.Sprint(offset),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.All(offset)
	return
}

// Get ...
func (mw *LoggingMiddleware) Get(id int) (output models.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"id", fmt.Sprint(id),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Get(id)
	return
}
