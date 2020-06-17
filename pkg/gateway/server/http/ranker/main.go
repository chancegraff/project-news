package ranker

import (
	web "net/http"

	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Endpoints ...
type Endpoints struct {
	ArticlesEndpoint *httptransport.Server
	UserEndpoint     *httptransport.Server
	VoteEndpoint     *httptransport.Server
}

// NewEndpoints ...
func NewEndpoints(endpoints endpoints.Endpoints) Endpoints {
	return Endpoints{
		ArticlesEndpoint: MakeArticlesEndpoint(&endpoints),
		UserEndpoint:     MakeUserEndpoint(&endpoints),
		VoteEndpoint:     MakeVoteEndpoint(&endpoints),
	}
}

// Route ...
func (e *Endpoints) Route(mxr *mux.Router) {
	route := mxr.PathPrefix("/ranker").Subrouter()
	route.HandleFunc("/articles", e.Articles).Methods("POST", "OPTIONS")
	route.HandleFunc("/user", e.User).Methods("POST", "OPTIONS")
	route.HandleFunc("/vote", e.Vote).Methods("POST", "OPTIONS")
}

// Articles ...
func (e *Endpoints) Articles(writer web.ResponseWriter, request *web.Request) {
	e.ArticlesEndpoint.ServeHTTP(writer, request)
}

// User ...
func (e *Endpoints) User(writer web.ResponseWriter, request *web.Request) {
	e.UserEndpoint.ServeHTTP(writer, request)
}

// Vote ...
func (e *Endpoints) Vote(writer web.ResponseWriter, request *web.Request) {
	e.VoteEndpoint.ServeHTTP(writer, request)
}
