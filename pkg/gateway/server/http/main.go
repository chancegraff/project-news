package http

import (
	"net/http"
	"path/filepath"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/server/http/auth"
	"github.com/chancegraff/project-news/pkg/gateway/server/http/collector"
	"github.com/chancegraff/project-news/pkg/gateway/server/http/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/server/http/token"

	"github.com/gorilla/mux"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	AuthEndpoints      auth.Endpoints
	CollectorEndpoints collector.Endpoints
	RankerEndpoints    ranker.Endpoints
	TokenEndpoints     token.Endpoints
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) ServerEndpoints {
	return ServerEndpoints{
		AuthEndpoints:      auth.NewEndpoints(endpoints),
		CollectorEndpoints: collector.NewEndpoints(endpoints),
		RankerEndpoints:    ranker.NewEndpoints(endpoints),
		TokenEndpoints:     token.NewEndpoints(endpoints),
	}
}

// Route will create a muxed server
func (e *ServerEndpoints) Route() *mux.Router {
	mxr := mux.NewRouter()

	// Create api
	api := mxr.PathPrefix("/api/v1").Subrouter()
	e.AuthEndpoints.Route(api)
	e.CollectorEndpoints.Route(api)
	e.RankerEndpoints.Route(api)
	e.TokenEndpoints.Route(api)

	// Create file server at root
	wd := utils.Getwd()
	fp := filepath.Join(wd, "web", "build")
	fs := http.FileServer(http.Dir(fp))
	mxr.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// TODO Setup webpack for golang if it's dev

	// // Create root for servers
	// env := utils.GetEnvironment()
	// if env != "dev" {
	// 	wd := utils.Getwd()
	// 	fp := filepath.Join(wd, "web", "build")
	// 	fs := http.FileServer(http.Dir(fp))
	// 	mxr.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	// }

	return mxr
}
