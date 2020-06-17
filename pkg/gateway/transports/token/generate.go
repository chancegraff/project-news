package token

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pbt "github.com/chancegraff/project-news/api/token"
)

// DecodeGenerateRequest ...
func DecodeGenerateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pbt.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeGenerateResponse ...
func EncodeGenerateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
