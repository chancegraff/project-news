package server

import (
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxGet ...
func (h *HTTP) MuxGet(mux *http.ServeMux) *http.ServeMux {
	get := httptransport.NewServer(
		h.endpoints.GetEndpoint,
		transports.DecodeGetRequest,
		transports.EncodeResponse,
	)
	mux.Handle("/get", get)
	return mux
}
