package rpc

import (
	"github.com/chancegraff/project-news/pkg/services/auth/endpoints"
	"github.com/chancegraff/project-news/pkg/services/auth/transports"
	gt "github.com/go-kit/kit/transport/grpc"
)

// MakeDeregisterEndpoint ...
func MakeDeregisterEndpoint(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.DeregisterEndpoint,
		transports.DecodeDeregisterRPCRequest,
		transports.EncodeDeregisterRPCResponse,
	)
}
