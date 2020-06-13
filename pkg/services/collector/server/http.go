package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// HTTP ...
type HTTP struct{}

// Start will begin the HTTP server
func (HTTP) Start(port string) error {
	return http.ListenAndServe(port, nil)
}

// NewHTTPServer ...
func NewHTTPServer(endpoints endpoints.Endpoints) HTTP {
	all := httptransport.NewServer(
		endpoints.All,
		transports.DecodeAllRequest,
		transports.EncodeResponse,
	)

	get := httptransport.NewServer(
		endpoints.Get,
		transports.DecodeGetRequest,
		transports.EncodeResponse,
	)

	http.Handle("/all", all)
	http.Handle("/get", get)

	return HTTP{}
}
