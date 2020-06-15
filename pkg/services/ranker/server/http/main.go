package http

import (
	web "net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

// ServerEndpoints ...
type ServerEndpoints struct {
	ArticlesEndpoint *httptransport.Server
	UserEndpoint     *httptransport.Server
	VoteEndpoint     *httptransport.Server
}

// NewServerEndpoints ...
func NewServerEndpoints(endpoints endpoints.Endpoints) ServerEndpoints {
	return ServerEndpoints{
		ArticlesEndpoint: MakeArticlesEndpoint(&endpoints),
		UserEndpoint:     MakeUserEndpoint(&endpoints),
		VoteEndpoint:     MakeVoteEndpoint(&endpoints),
	}
}

// Articles ...
func (s *ServerEndpoints) Articles(writer web.ResponseWriter, request *web.Request) {
	s.ArticlesEndpoint.ServeHTTP(writer, request)
}

// User ...
func (s *ServerEndpoints) User(writer web.ResponseWriter, request *web.Request) {
	s.UserEndpoint.ServeHTTP(writer, request)
}

// Vote ...
func (s *ServerEndpoints) Vote(writer web.ResponseWriter, request *web.Request) {
	s.VoteEndpoint.ServeHTTP(writer, request)
}
