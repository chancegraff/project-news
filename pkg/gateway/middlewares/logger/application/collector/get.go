package collector

import (
	"fmt"
	"time"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// Get ...
func (mw *Middleware) Get(id int) (output *pbc.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Get",
			"id", fmt.Sprint(id),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Get(id)
	return
}
