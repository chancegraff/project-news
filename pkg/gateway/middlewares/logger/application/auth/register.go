package auth

import (
	"fmt"
	"time"

	pba "github.com/chancegraff/project-news/api/auth"
)

// Register ...
func (mw *Middleware) Register(email string, password string) (output *pba.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Register",
			"email", fmt.Sprint(email),
			"output", fmt.Sprint(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Register(email, password)
	return
}
