package token

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pbt "github.com/chancegraff/project-news/api/token"
)

// DecodeVerifyRequest ...
func DecodeVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pbt.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeVerifyResponse ...
func EncodeVerifyResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
