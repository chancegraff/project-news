package routes

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	gt "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

// VoteHTTP ...
func VoteHTTP(endpoints *endpoints.Endpoints, mux *http.ServeMux) *http.ServeMux {
	vote := httptransport.NewServer(
		endpoints.VoteEndpoint,
		transports.DecodeVoteHTTPRequest,
		transports.EncodeHTTPResponse,
	)
	mux.Handle("/vote", vote)
	return mux
}

// VoteRPC ...
func VoteRPC(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.VoteEndpoint,
		transports.DecodeVoteRPCRequest,
		transports.EncodeVoteRPCResponse,
	)
}
