package token

import (
	web "net/http"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Endpoints ...
type Endpoints struct {
	GenerateEndpoint *httptransport.Server
	VerifyEndpoint   *httptransport.Server
}

// NewEndpoints ...
func NewEndpoints(endpoints endpoints.Endpoints) Endpoints {
	return Endpoints{
		GenerateEndpoint: MakeGenerateEndpoint(&endpoints),
		VerifyEndpoint:   MakeVerifyEndpoint(&endpoints),
	}
}

// Route ...
func (e *Endpoints) Route(mxr *mux.Router) {
	route := mxr.PathPrefix("/token").Subrouter()
	route.HandleFunc("/generate", e.Generate).Methods("POST", "OPTIONS")
	route.HandleFunc("/verify", e.Verify).Methods("POST", "OPTIONS")
}

// Generate ...
func (e *Endpoints) Generate(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.GenerateEndpoint.ServeHTTP(writer, request)
}

// Verify ...
func (e *Endpoints) Verify(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.VerifyEndpoint.ServeHTTP(writer, request)
}
