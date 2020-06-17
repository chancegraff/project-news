package collector

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// AllRequest ...
type AllRequest struct {
	Offset int `json:"offset,omitempty"`
}

// DecodeAllRequest ...
func DecodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pbc.AllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeAllResponse ...
func EncodeAllResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
