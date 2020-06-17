package auth

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	transports "github.com/chancegraff/project-news/pkg/gateway/transports/auth"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.AuthEndpoints.VerifyEndpoint,
		transports.DecodeVerifyRequest,
		transports.EncodeVerifyResponse,
	)
}
