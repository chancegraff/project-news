package routes

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	gt "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

// ArticlesHTTP ...
func ArticlesHTTP(endpoints *endpoints.Endpoints, mux *http.ServeMux) *http.ServeMux {
	articles := httptransport.NewServer(
		endpoints.ArticlesEndpoint,
		transports.DecodeArticlesHTTPRequest,
		transports.EncodeHTTPResponse,
	)
	mux.Handle("/articles", articles)
	return mux
}

// ArticlesRPC ...
func ArticlesRPC(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.ArticlesEndpoint,
		transports.DecodeArticlesRPCRequest,
		transports.EncodeArticlesRPCResponse,
	)
}
