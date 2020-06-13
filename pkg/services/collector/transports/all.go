package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/interfaces"
)

// DecodeAllRequest ...
func DecodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request interfaces.AllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
