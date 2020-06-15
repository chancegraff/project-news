package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/server/routes"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// HTTP ...
type HTTP struct {
	endpoints *endpoints.Endpoints
	server    *http.Server
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

// NewMux will create a muxer with the routes registered
func (h *HTTP) NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux = routes.ArticlesHTTP(h.endpoints, mux)
	mux = routes.UserHTTP(h.endpoints, mux)
	mux = routes.VoteHTTP(h.endpoints, mux)
	return mux
}

// NewHTTPServer instantiates a new HTTP server with the services endpoints
func NewHTTPServer(endpoints endpoints.Endpoints) *HTTP {
	// Create the address
	port := utils.GetRankerPort()
	address := fmt.Sprint(":", port)

	// Create HTTP from file
	h := HTTP{
		endpoints: &endpoints,
		port:      port,
		address:   address,
	}

	// Create Server from library
	h.server = &http.Server{
		Handler:      h.NewMux(),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Return HTTP
	return &h
}
