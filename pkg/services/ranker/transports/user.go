package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// UserRequest ...
type UserRequest struct {
	UserID string `json:"user"`
}

// UserResponse ...
type UserResponse struct {
	User models.UserVotes `json:"user"`
	Err  string           `json:"err,omitempty"`
}

// DecodeUserRequest ...
func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}
