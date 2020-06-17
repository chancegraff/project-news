package ranker

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	transports "github.com/chancegraff/project-news/pkg/gateway/transports/ranker"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeArticlesEndpoint ...
func MakeArticlesEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.RankerEndpoints.ArticlesEndpoint,
		transports.DecodeArticlesRequest,
		transports.EncodeArticlesResponse,
	)
}
