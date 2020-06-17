package rpc

import (
	"context"

	pb "github.com/chancegraff/project-news/api/auth"
	"github.com/chancegraff/project-news/pkg/services/auth/endpoints"
	gt "github.com/go-kit/kit/transport/grpc"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	DeregisterEndpoint gt.Handler
	RegisterEndpoint   gt.Handler
	UserEndpoint       gt.Handler
	VerifyEndpoint     gt.Handler
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) pb.AuthServiceServer {
	return &ServerEndpoints{
		DeregisterEndpoint: MakeDeregisterEndpoint(&endpoints),
		RegisterEndpoint:   MakeRegisterEndpoint(&endpoints),
		UserEndpoint:       MakeUserEndpoint(&endpoints),
		VerifyEndpoint:     MakeVerifyEndpoint(&endpoints),
	}
}

// Deregister ...
func (s *ServerEndpoints) Deregister(ctx context.Context, req *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	_, resp, err := s.DeregisterEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeregisterResponse), nil
}

// Register ...
func (s *ServerEndpoints) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, resp, err := s.RegisterEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RegisterResponse), nil
}

// User ...
func (s *ServerEndpoints) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.UserEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

// Verify ...
func (s *ServerEndpoints) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	_, resp, err := s.VerifyEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VerifyResponse), nil
}
