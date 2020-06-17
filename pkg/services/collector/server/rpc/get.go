package rpc

import (
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	gt "github.com/go-kit/kit/transport/grpc"
)

// MakeGetEndpoint ...
func MakeGetEndpoint(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.GetEndpoint,
		transports.DecodeGetRPCRequest,
		transports.EncodeGetRPCResponse,
	)
}
