package rpc

import (
	"context"

	pb "github.com/chancegraff/project-news/api/collector"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	gt "github.com/go-kit/kit/transport/grpc"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	AllEndpoint gt.Handler
	GetEndpoint gt.Handler
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) pb.CollectorServiceServer {
	return &ServerEndpoints{
		AllEndpoint: MakeAllEndpoint(&endpoints),
		GetEndpoint: MakeGetEndpoint(&endpoints),
	}
}

// All ...
func (s *ServerEndpoints) All(ctx context.Context, req *pb.AllRequest) (*pb.AllResponse, error) {
	_, resp, err := s.AllEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AllResponse), nil
}

// Get ...
func (s *ServerEndpoints) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := s.GetEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}
