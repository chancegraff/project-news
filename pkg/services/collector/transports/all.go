package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// AllRequest ...
type AllRequest struct {
	Offset int `json:"offset,omitempty"`
}

// AllResponse ...
type AllResponse struct {
	Articles []models.Article `json:"articles"`
	Err      string           `json:"err,omitempty"`
}

// DecodeAllRequest ...
func DecodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}
