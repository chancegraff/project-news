package http

import (
	web "net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MuxAll ...
func (h *HTTP) MuxAll(mux *web.ServeMux) *web.ServeMux {
	all := httptransport.NewServer(
		h.endpoints.AllEndpoint,
		transports.DecodeAllHTTPRequest,
		transports.EncodeHTTPResponse,
	)
	mux.Handle("/all", all)
	return mux
}
