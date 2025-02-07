package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pba "github.com/chancegraff/project-news/api/auth"
)

// DecodeVerifyRequest ...
func DecodeVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pba.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeVerifyResponse ...
func EncodeVerifyResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
