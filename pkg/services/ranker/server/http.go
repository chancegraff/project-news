package server

import (
	"context"
	"fmt"
	web "net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/server/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// HTTP ...
type HTTP struct {
	endpoints http.ServerEndpoints
	server    *web.Server
	address   string
	port      int
}

// Start will begin the HTTP server
func (h *HTTP) Start(parent context.Context, logger log.Logger) error {
	_, cancel := context.WithCancel(parent)
	level.Info(logger).Log("msg", "service started")
	err := h.server.ListenAndServe()
	cancel()
	return err
}

// Stop will stop the HTTP server
func (h *HTTP) Stop(parent context.Context, logger log.Logger) error {
	level.Info(logger).Log("msg", "service stopped")
	return h.server.Shutdown(parent)
}

// Mux will register the routes to a muxer and return it
func Mux(endpoints http.ServerEndpoints) *web.ServeMux {
	mux := web.NewServeMux()
	mux.HandleFunc("/articles", endpoints.Articles)
	mux.HandleFunc("/user", endpoints.User)
	mux.HandleFunc("/vote", endpoints.Vote)
	return mux
}

// NewServer will create a new HTTP server
func NewServer(e endpoints.Endpoints) HTTP {
	// Create the address
	port := utils.GetRankerPort()
	address := fmt.Sprint(":", port)

	// Build the endpoints
	endpoints := http.NewServerEndpoints(e)
	handler := Mux(endpoints)

	// Bind a listener
	server := &web.Server{
		Handler:      handler,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Return HTTP interface
	return HTTP{
		server:    server,
		endpoints: endpoints,
		address:   address,
		port:      port,
	}
}
