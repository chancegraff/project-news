package token

import (
	"fmt"
	"time"

	pbt "github.com/chancegraff/project-news/api/token"
)

// Generate ...
func (mw *Middleware) Generate(identifiers *pbt.Identifiers, client *pbt.Client) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Generate",
			"identifiers", fmt.Sprint(identifiers),
			"client", fmt.Sprint(client),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Generate(identifiers, client)
	return
}
