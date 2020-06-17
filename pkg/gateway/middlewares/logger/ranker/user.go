package ranker

import (
	"fmt"
	"time"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// User ...
func (mw *LoggingMiddleware) User(userID string) (output *pbr.UserVotes, err error) {
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
