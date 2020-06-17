package ranker

import (
	"fmt"
	"time"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Articles ...
func (mw *LoggingMiddleware) Articles(articleIDs []string) (output []*pbr.ArticleVotes, err error) {
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
