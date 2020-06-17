package rpc

import (
	"context"

	pb "github.com/chancegraff/project-news/api/token"
	"github.com/chancegraff/project-news/pkg/services/token/endpoints"
	gt "github.com/go-kit/kit/transport/grpc"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	GenerateEndpoint gt.Handler
	VerifyEndpoint   gt.Handler
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) pb.TokenServiceServer {
	return &ServerEndpoints{
		GenerateEndpoint: MakeGenerateEndpoint(&endpoints),
		VerifyEndpoint:   MakeVerifyEndpoint(&endpoints),
	}
}

// Generate ...
func (s *ServerEndpoints) Generate(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	_, resp, err := s.GenerateEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GenerateResponse), nil
}

// Verify ...
func (s *ServerEndpoints) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	_, resp, err := s.VerifyEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VerifyResponse), nil
}
