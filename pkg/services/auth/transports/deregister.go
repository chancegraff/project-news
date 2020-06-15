package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// DeregisterRequest ...
type DeregisterRequest struct {
	UserID string `json:"user"`
}

// DeregisterResponse ...
type DeregisterResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeDeregisterRequest ...
func DecodeDeregisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeregisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}
