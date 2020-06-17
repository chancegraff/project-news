package collector

import (
	"fmt"
	"time"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// All ...
func (mw *LoggingMiddleware) All(offset int) (output []*pbc.Article, err error) {
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
