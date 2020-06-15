package rpc

import (
	"context"

	pb "github.com/chancegraff/project-news/api/ranker"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	gt "github.com/go-kit/kit/transport/grpc"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	ArticlesEndpoint gt.Handler
	UserEndpoint     gt.Handler
	VoteEndpoint     gt.Handler
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) pb.RankerServiceServer {
	return &ServerEndpoints{
		ArticlesEndpoint: MakeArticlesEndpoint(&endpoints),
		UserEndpoint:     MakeUserEndpoint(&endpoints),
		VoteEndpoint:     MakeVoteEndpoint(&endpoints),
	}
}

// Articles ...
func (s *ServerEndpoints) Articles(ctx context.Context, req *pb.ArticlesRequest) (*pb.ArticlesResponse, error) {
	_, resp, err := s.ArticlesEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArticlesResponse), nil
}

// User ...
func (s *ServerEndpoints) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.UserEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

// Vote ...
func (s *ServerEndpoints) Vote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	_, resp, err := s.VoteEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VoteResponse), nil
}
