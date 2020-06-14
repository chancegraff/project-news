package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxUser ...
func (h *HTTP) MuxUser(mux *http.ServeMux) *http.ServeMux {
	user := httptransport.NewServer(
		h.endpoints.UserEndpoint,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/user", user)
	return mux
}
