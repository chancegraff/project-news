package http

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeArticlesEndpoint ...
func MakeArticlesEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.ArticlesEndpoint,
		transports.DecodeArticlesHTTPRequest,
		transports.EncodeHTTPResponse,
	)
}
