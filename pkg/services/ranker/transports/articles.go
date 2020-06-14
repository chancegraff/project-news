package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// ArticlesRequest ...
type ArticlesRequest struct {
	ArticleIDs []string `json:"articles"`
}

// ArticlesResponse ...
type ArticlesResponse struct {
	Articles []models.ArticleVotes `json:"articles"`
	Err      string                `json:"err,omitempty"`
}

// DecodeArticlesRequest ...
func DecodeArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ArticlesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
