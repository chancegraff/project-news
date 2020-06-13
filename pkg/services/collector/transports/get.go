package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// GetRequest ...
type GetRequest struct {
	ID int `json:"id"`
}

// GetResponse ...
type GetResponse struct {
	Article models.Article `json:"article"`
	Err     string         `json:"err,omitempty"`
}

// DecodeGetRequest ...
func DecodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
