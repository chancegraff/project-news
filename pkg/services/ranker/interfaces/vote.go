package interfaces

import "github.com/chancegraff/project-news/internal/models"

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
