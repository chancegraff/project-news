package collector

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// DecodeGetRequest ...
func DecodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pbc.GetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeGetResponse ...
func EncodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
