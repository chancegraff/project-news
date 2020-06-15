package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// VerifyRequest ...
type VerifyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// VerifyResponse ...
type VerifyResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeVerifyRequest ...
func DecodeVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}
