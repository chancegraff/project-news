package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pba "github.com/chancegraff/project-news/api/auth"
)

// DecodeDeregisterRequest ...
func DecodeDeregisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request *pba.DeregisterRequest
	if err := json.NewDecoder(r.Body).Decode(request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeDeregisterResponse ...
func EncodeDeregisterResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
