package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pba "github.com/chancegraff/project-news/api/auth"
)

// DecodeUserRequest ...
func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pba.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeUserResponse ...
func EncodeUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
