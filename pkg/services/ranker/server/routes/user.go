package routes

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	gt "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

// UserHTTP ...
func UserHTTP(endpoints *endpoints.Endpoints, mux *http.ServeMux) *http.ServeMux {
	user := httptransport.NewServer(
		endpoints.UserEndpoint,
		transports.DecodeUserHTTPRequest,
		transports.EncodeHTTPResponse,
	)
	mux.Handle("/user", user)
	return mux
}

// UserRPC ...
func UserRPC(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.UserEndpoint,
		transports.DecodeUserRPCRequest,
		transports.EncodeUserRPCResponse,
	)
}
