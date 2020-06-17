package token

import (
	"fmt"
	"time"

	pbt "github.com/chancegraff/project-news/api/token"
)

// Verify ...
func (mw *Middleware) Verify(identifiers *pbt.Identifiers, client *pbt.Client) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Verify",
			"identifiers", fmt.Sprint(identifiers),
			"client", fmt.Sprint(client),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Verify(identifiers, client)
	return
}
