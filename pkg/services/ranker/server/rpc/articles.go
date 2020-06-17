package rpc

import (
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	gt "github.com/go-kit/kit/transport/grpc"
)

// MakeArticlesEndpoint ...
func MakeArticlesEndpoint(endpoints *endpoints.Endpoints) *gt.Server {
	return gt.NewServer(
		endpoints.ArticlesEndpoint,
		transports.DecodeArticlesRPCRequest,
		transports.EncodeArticlesRPCResponse,
	)
}
