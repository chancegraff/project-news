package transport

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/auth"
	"github.com/go-kit/kit/log"
)

// MakeAuthMiddleware ...
func MakeAuthMiddleware(logger log.Logger) auth.Middleware {
	lgr := log.With(logger, "end", "auth")
	return func(next auth.Endpoints) auth.Endpoints {
		return auth.Endpoints{
			DeregisterEndpoint: MakeEndpoint("Deregister", lgr, next.DeregisterEndpoint),
			RegisterEndpoint:   MakeEndpoint("Register", lgr, next.RegisterEndpoint),
			UserEndpoint:       MakeEndpoint("User", lgr, next.UserEndpoint),
			VerifyEndpoint:     MakeEndpoint("Verify", lgr, next.VerifyEndpoint),
		}
	}
}
