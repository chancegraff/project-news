package auth

import (
	web "net/http"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Endpoints ...
type Endpoints struct {
	DeregisterEndpoint *httptransport.Server
	RegisterEndpoint   *httptransport.Server
	UserEndpoint       *httptransport.Server
	VerifyEndpoint     *httptransport.Server
}

// NewEndpoints ...
func NewEndpoints(endpoints endpoints.Endpoints) Endpoints {
	return Endpoints{
		DeregisterEndpoint: MakeDeregisterEndpoint(&endpoints),
		RegisterEndpoint:   MakeRegisterEndpoint(&endpoints),
		UserEndpoint:       MakeUserEndpoint(&endpoints),
		VerifyEndpoint:     MakeVerifyEndpoint(&endpoints),
	}
}

// Route ...
func (e *Endpoints) Route(mxr *mux.Router) {
	route := mxr.PathPrefix("/auth").Subrouter()
	route.HandleFunc("/deregister", e.Deregister).Methods("POST", "OPTIONS")
	route.HandleFunc("/register", e.Register).Methods("POST", "OPTIONS")
	route.HandleFunc("/user", e.User).Methods("POST", "OPTIONS")
	route.HandleFunc("/verify", e.Verify).Methods("POST", "OPTIONS")
}

// Deregister ...
func (e *Endpoints) Deregister(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.DeregisterEndpoint.ServeHTTP(writer, request)
}

// Register ...
func (e *Endpoints) Register(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.RegisterEndpoint.ServeHTTP(writer, request)
}

// User ...
func (e *Endpoints) User(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.UserEndpoint.ServeHTTP(writer, request)
}

// Verify ...
func (e *Endpoints) Verify(writer web.ResponseWriter, request *web.Request) {
	utils.SetCORSHeaders(writer)

	if request.Method == "OPTIONS" {
		return
	}

	e.VerifyEndpoint.ServeHTTP(writer, request)
}
