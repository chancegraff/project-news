package collector

import (
	"context"
	"fmt"
	"time"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// All ...
func (mw *Middleware) All(ctx context.Context, offset int) (output []*pbc.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "All",
			"offset", fmt.Sprint(offset),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.All(ctx, offset)
	return
}
