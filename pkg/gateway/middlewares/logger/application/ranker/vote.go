package ranker

import (
	"fmt"
	"time"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Vote ...
func (mw *Middleware) Vote(articleID, userID string) (output *pbr.ArticleVotes, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Vote",
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
