package middlewares

import (
	"fmt"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
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

// Deregister ...
func (mw *LoggingMiddleware) Deregister(userID string) (output models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"userID", fmt.Sprint(userID),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Deregister(userID)
	return
}

// Register ...
func (mw *LoggingMiddleware) Register(email string, password string) (output models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"email", fmt.Sprint(email),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Register(email, password)
	return
}

// User ...
func (mw *LoggingMiddleware) User(userID string) (output models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"userID", fmt.Sprint(userID),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.User(userID)
	return
}

// Verify ...
func (mw *LoggingMiddleware) Verify(email string, password string) (output models.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"email", fmt.Sprint(email),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Verify(email, password)
	return
}
