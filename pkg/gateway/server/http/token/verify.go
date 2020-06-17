package token

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	transports "github.com/chancegraff/project-news/pkg/gateway/transports/token"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.TokenEndpoints.VerifyEndpoint,
		transports.DecodeVerifyRequest,
		transports.EncodeVerifyResponse,
	)
}
