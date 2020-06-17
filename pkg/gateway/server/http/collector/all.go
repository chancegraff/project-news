package collector

import (
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	transports "github.com/chancegraff/project-news/pkg/gateway/transports/collector"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeAllEndpoint ...
func MakeAllEndpoint(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.CollectorEndpoints.AllEndpoint,
		transports.DecodeAllRequest,
		transports.EncodeAllResponse,
	)
}
