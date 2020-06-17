package transport

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints/ranker"
	"github.com/go-kit/kit/log"
)

// MakeRankerMiddleware ...
func MakeRankerMiddleware(logger log.Logger) ranker.Middleware {
	lgr := log.With(logger, "end", "ranker")
	return func(next ranker.Endpoints) ranker.Endpoints {
		return ranker.Endpoints{
			ArticlesEndpoint: MakeEndpoint("Articles", lgr, next.ArticlesEndpoint),
			UserEndpoint:     MakeEndpoint("User", lgr, next.UserEndpoint),
			VoteEndpoint:     MakeEndpoint("Vote", lgr, next.VoteEndpoint),
		}
	}
}
