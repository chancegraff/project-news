package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
)

// HTTP ...
type HTTP struct {
	server  *http.Server
	address string
	port    int
}

// Start will begin the HTTP server
func (h HTTP) Start() error {
	log.Printf("Collector started at %s", h.address)
	return h.server.ListenAndServe()
}

// Stop will shut down the HTTP server
func (h HTTP) Stop() error {
	log.Println("Collector stopped")
	return h.server.Close()
}

var getCORS = handlers.CORS(
	handlers.AllowedHeaders(
		[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
	),
	handlers.AllowedMethods(
		[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	),
)

// NewHTTPServer ...
func NewHTTPServer(endpoints endpoints.Endpoints) HTTP {
	port := utils.GetCollectorPort()
	address := fmt.Sprint(":", port)
	mux := http.NewServeMux()

	all := httptransport.NewServer(
		endpoints.All,
		transports.DecodeAllRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/all", all)

	get := httptransport.NewServer(
		endpoints.Get,
		transports.DecodeGetRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/get", get)

	server := http.Server{
		Handler:      getCORS(mux),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return HTTP{&server, address, port}
}
