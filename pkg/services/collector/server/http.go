package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
)

// HTTP ...
type HTTP struct {
	endpoints *endpoints.Endpoints
	server    *http.Server
	address   string
	port      int
}

// Start will begin the HTTP server
func (h *HTTP) Start(parent context.Context) error {
	_, cancel := context.WithCancel(parent)
	log.Printf("Server started at %s", h.address)
	err := h.server.ListenAndServe()
	cancel()
	return err
}

// Stop will shut down the HTTP server
func (h *HTTP) Stop(parent context.Context) error {
	log.Printf("HTTP stopped")
	return h.server.Shutdown(parent)
}

// MuxRoutes will create a muxer with routes registered
func (h *HTTP) MuxRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux = h.MuxAll(mux)
	mux = h.MuxGet(mux)
	return mux
}

// NewHTTPServer ...
func NewHTTPServer(endpoints endpoints.Endpoints) *HTTP {
	port := utils.GetCollectorPort()
	address := fmt.Sprint(":", port)
	h := HTTP{
		endpoints: &endpoints,
		port:      port,
		address:   address,
	}

	h.server = &http.Server{
		Handler:      h.MuxRoutes(),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &h
}
