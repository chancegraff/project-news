package middlewares

import (
	"fmt"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
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

// Articles ...
func (mw *LoggingMiddleware) Articles(articleIDs []string) (output []models.ArticleVotes, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"articleIDs", fmt.Sprint(articleIDs),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Articles(articleIDs)
	return
}

// User ...
func (mw *LoggingMiddleware) User(userID string) (output models.UserVotes, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "user",
			"userID", fmt.Sprint(userID),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.User(userID)
	return
}

// Vote ...
func (mw *LoggingMiddleware) Vote(articleID, userID string) (output models.ArticleVotes, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "ranker",
			"method", "articles",
			"articleID", fmt.Sprint(articleID),
			"userID", fmt.Sprint(userID),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Vote(articleID, userID)
	return
}
