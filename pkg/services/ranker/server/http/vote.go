package http

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeVoteEndpoint ...
func MakeVoteEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.VoteEndpoint,
		transports.DecodeVoteHTTPRequest,
		transports.EncodeHTTPResponse,
	)
}
