package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// HTTP ...
type HTTP struct{}

// NewHTTPServer instantiates a new HTTP server with the services endpoints
func NewHTTPServer(endpoints endpoints.Endpoints) HTTP {
	articles := httptransport.NewServer(
		endpoints.Articles,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)
	user := httptransport.NewServer(
		endpoints.User,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)
	vote := httptransport.NewServer(
		endpoints.Vote,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)

	http.Handle("/articles", articles)
	http.Handle("/user", user)
	http.Handle("/vote", vote)

	return HTTP{}
}

// Start will begin the HTTP server
func (HTTP) Start(port string) error {
	return http.ListenAndServe(port, nil)
}
