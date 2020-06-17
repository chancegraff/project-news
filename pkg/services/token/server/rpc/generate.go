package rpc

import (
	"github.com/chancegraff/project-news/pkg/services/token/endpoints"
	"github.com/chancegraff/project-news/pkg/services/token/transports"
	gt "github.com/go-kit/kit/transport/grpc"
)

// MakeGenerateEndpoint ...
func MakeGenerateEndpoint(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.GenerateEndpoint,
		transports.DecodeGenerateRPCRequest,
		transports.EncodeGenerateRPCResponse,
	)
}
