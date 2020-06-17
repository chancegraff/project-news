package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/chancegraff/project-news/api/collector"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/server/rpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

// RPC ...
type RPC struct {
	endpoints pb.CollectorServiceServer
	server    *grpc.Server
	listener  net.Listener
	address   string
	port      int
}

// Start ...
func (r *RPC) Start(parent context.Context, logger log.Logger) error {
	_, cancel := context.WithCancel(parent)
	level.Info(logger).Log("msg", "service started")
	err := r.server.Serve(r.listener)
	cancel()
	return err
}

// Stop ...
func (r *RPC) Stop(parent context.Context, logger log.Logger) error {
	level.Info(logger).Log("msg", "service stopped")
	err := r.listener.Close()
	if err != nil {
		return err
	}
	r.server.GracefulStop()
	return nil
}

// NewRPCServer ...
func NewRPCServer(e endpoints.Endpoints) RPC {
	// Create the address
	port := utils.GetCollectorPort()
	address := fmt.Sprint(":", port)

	// Bind a listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	// Bind to protobuffs
	endpoints := rpc.NewServerEndpoints(e)
	server := grpc.NewServer()
	pb.RegisterCollectorServiceServer(server, endpoints)

	// Return RPC interface
	return RPC{
		endpoints: endpoints,
		server:    server,
		listener:  listener,
		port:      port,
		address:   address,
	}
}
