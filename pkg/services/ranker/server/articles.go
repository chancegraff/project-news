package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxArticles ...
func (h *HTTP) MuxArticles(mux *http.ServeMux) *http.ServeMux {
	articles := httptransport.NewServer(
		h.endpoints.ArticlesEndpoint,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/articles", articles)
	return mux
}
