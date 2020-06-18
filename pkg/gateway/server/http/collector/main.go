package collector

import (
	web "net/http"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Endpoints ...
type Endpoints struct {
	AllEndpoint *httptransport.Server
	GetEndpoint *httptransport.Server
}

// NewEndpoints ...
func NewEndpoints(endpoints endpoints.Endpoints) Endpoints {
	return Endpoints{
		AllEndpoint: MakeAllEndpoint(&endpoints),
		GetEndpoint: MakeGetEndpoint(&endpoints),
	}
}

// Route ...
func (e *Endpoints) Route(mxr *mux.Router) {
	route := mxr.PathPrefix("/collector").Subrouter()
	route.HandleFunc("/all", e.All).Methods("POST", "OPTIONS")
	route.HandleFunc("/get", e.Get).Methods("POST", "OPTIONS")
}

// All ...
func (e *Endpoints) All(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.AllEndpoint.ServeHTTP(writer, request)
}

// Get ...
func (e *Endpoints) Get(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.GetEndpoint.ServeHTTP(writer, request)
}
