package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/chancegraff/project-news/api/ranker"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/server/rpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

// RPC ...
type RPC struct {
	endpoints *endpoints.Endpoints
	server    pb.RankerServiceServer
	listener  net.Listener
	address   string
	port      int
}

// Start ...
func (r *RPC) Start(parent context.Context, logger log.Logger) error {
	_, cancel := context.WithCancel(parent)
	listener, err := net.Listen("tcp", r.address)
	if err != nil {
		cancel()
		return err
	}
	r.listener = listener
	server := grpc.NewServer()
	level.Info(logger).Log("msg", "service started")
	pb.RegisterRankerServiceServer(server, r.server)
	err = server.Serve(r.listener)
	cancel()
	return err
}

// Stop ...
func (r *RPC) Stop(parent context.Context, logger log.Logger) error {
	level.Info(logger).Log("msg", "service stopped")
	return r.listener.Close()
}

// NewRPCServer ...
func NewRPCServer(endpoints endpoints.Endpoints) RPC {
	// Create the address
	port := utils.GetRankerPort()
	address := fmt.Sprint(":", port)

	// Return RPC interface
	return RPC{
		endpoints: &endpoints,
		server:    rpc.NewServer(endpoints),
		port:      port,
		address:   address,
	}
}
