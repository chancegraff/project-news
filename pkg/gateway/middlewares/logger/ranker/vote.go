package ranker

import (
	"fmt"
	"time"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Vote ...
func (mw *LoggingMiddleware) Vote(articleID, userID string) (output *pbr.ArticleVotes, err error) {
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
