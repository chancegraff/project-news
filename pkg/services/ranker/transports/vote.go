package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"
)

// VoteRequest ...
type VoteRequest struct {
	ArticleID string `json:"article"`
	UserID    string `json:"user"`
}

// VoteResponse ...
type VoteResponse struct {
	Article models.ArticleVotes `json:"article"`
	Err     string              `json:"err,omitempty"`
}

// DecodeVoteRequest ...
func DecodeVoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}
