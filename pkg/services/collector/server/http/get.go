package http

import (
	web "net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxGet ...
func (h *HTTP) MuxGet(mux *web.ServeMux) *web.ServeMux {
	get := httptransport.NewServer(
		h.endpoints.GetEndpoint,
		transports.DecodeGetHTTPRequest,
		transports.EncodeHTTPResponse,
	)
	mux.Handle("/get", get)
	return mux
}
