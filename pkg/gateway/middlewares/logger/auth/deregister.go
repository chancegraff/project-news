package auth

import (
	"fmt"
	"time"

	pba "github.com/chancegraff/project-news/api/auth"
)

// Deregister ...
func (mw *LoggingMiddleware) Deregister(userID string) (output *pba.User, err error) {
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
