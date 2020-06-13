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

// DecodeArticlesResponse ...
func DecodeArticlesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response ArticlesResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
