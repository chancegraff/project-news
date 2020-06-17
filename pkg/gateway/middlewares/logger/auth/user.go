package auth

import (
	"fmt"
	"time"

	pba "github.com/chancegraff/project-news/api/auth"
)

// User ...
func (mw *LoggingMiddleware) User(userID string) (output *pba.User, err error) {
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
