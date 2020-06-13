package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// AllRequest ...
type AllRequest struct {
	Offset int `json:"offset"`
}

// AllResponse ...
type AllResponse struct {
	Articles []models.Article `json:"articles"`
	Err      string           `json:"err,omitempty"`
}

// DecodeAllRequest ...
func DecodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
