package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxVote ...
func (h *HTTP) MuxVote(mux *http.ServeMux) *http.ServeMux {
	vote := httptransport.NewServer(
		h.endpoints.VoteEndpoint,
		transports.DecodeArticlesRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/vote", vote)
	return mux
}
