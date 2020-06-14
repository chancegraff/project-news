package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxAll ...
func (h *HTTP) MuxAll(mux *http.ServeMux) *http.ServeMux {
	all := httptransport.NewServer(
		h.endpoints.AllEndpoint,
		transports.DecodeAllRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/all", all)
	return mux
}
