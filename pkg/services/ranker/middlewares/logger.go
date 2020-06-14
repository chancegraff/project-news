package middlewares

import (
	"os"
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
func MakeLoggingMiggleware() service.Middleware {
	return func(next service.Service) service.Service {
		logger := log.NewLogfmtLogger(os.Stderr)
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
			"method", "articles",
			"articleIDs", articleIDs,
			"output", output,
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
			"method", "user",
			"userID", userID,
			"output", output,
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
			"method", "articles",
			"articleID", articleID,
			"userID", userID,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Vote(articleID, userID)
	return
}
