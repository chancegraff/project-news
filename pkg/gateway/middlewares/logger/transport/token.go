package transport

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/token"
	"github.com/go-kit/kit/log"
)

// MakeTokenMiddleware ...
func MakeTokenMiddleware(logger log.Logger) token.Middleware {
	lgr := log.With(logger, "end", "token")
	return func(next token.Endpoints) token.Endpoints {
		return token.Endpoints{
			GenerateEndpoint: MakeEndpoint("Generate", lgr, next.GenerateEndpoint),
			VerifyEndpoint:   MakeEndpoint("Verify", lgr, next.VerifyEndpoint),
		}
	}
}
