package rpc

import (
	"context"

	pb "github.com/chancegraff/project-news/api/ranker"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/server/routes"
	gt "github.com/go-kit/kit/transport/grpc"
)

// Server ...
type Server struct {
	articles gt.Handler
	user     gt.Handler
	vote     gt.Handler
}

// NewServer ...
func NewServer(endpoints endpoints.Endpoints) pb.RankerServiceServer {
	return &Server{
		articles: routes.ArticlesRPC(&endpoints),
		user:     routes.UserRPC(&endpoints),
		vote:     routes.VoteRPC(&endpoints),
	}
}

// Articles ...
func (r *Server) Articles(ctx context.Context, req *pb.ArticlesRequest) (*pb.ArticlesResponse, error) {
	_, resp, err := r.articles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ArticlesResponse), nil
}

// User ...
func (r *Server) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := r.user.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

// Vote ...
func (r *Server) Vote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	_, resp, err := r.vote.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VoteResponse), nil
}
