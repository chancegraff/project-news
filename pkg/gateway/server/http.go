package server

import (
	"context"
	"fmt"
	web "net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/server/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// HTTP ...
type HTTP struct {
	endpoints http.ServerEndpoints
	server    *web.Server
	logger    log.Logger
	address   string
	port      int
}

// Start will begin the HTTP server
func (h *HTTP) Start(parent context.Context) error {
	_, cancel := context.WithCancel(parent)
	level.Info(h.logger).Log("msg", "service started", "address", h.address)
	err := h.server.ListenAndServe()
	cancel()
	return err
}

// Stop will stop the HTTP server
func (h *HTTP) Stop(parent context.Context) error {
	level.Info(h.logger).Log("msg", "service stopped")
	return h.server.Shutdown(parent)
}

// NewServer will create a new HTTP server
func NewServer(e endpoints.Endpoints, lgr log.Logger) HTTP {
	// Create the address
	port := utils.GetGatewayPort()
	address := fmt.Sprint(":", port)

	// Build the endpoints
	endpoints := http.NewServerEndpoints(e)
	handler := endpoints.Route()

	// handler.HandleFunc("/", func(rw web.ResponseWriter, rq *web.Request) {
	// 	handler.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 		tpl, err1 := route.GetPathTemplate()
	// 		met, err2 := route.GetMethods()
	// 		level.Info(lgr).Log("tpl", tpl, "err1", err1, "met", met, "err2", err2)
	// 		return nil
	// 	})
	// })

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
		logger:    lgr,
		endpoints: endpoints,
		address:   address,
		port:      port,
	}
}
