package http

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeUserEndpoint ...
func MakeUserEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.UserEndpoint,
		transports.DecodeUserHTTPRequest,
		transports.EncodeHTTPResponse,
	)
}
