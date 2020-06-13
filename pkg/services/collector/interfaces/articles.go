package interfaces

import "github.com/chancegraff/project-news/internal/models"

// ArticlesRequest ...
type ArticlesRequest struct {
	ArticleIDs []string `json:"articles"`
}

// ArticlesResponse ...
type ArticlesResponse struct {
	Articles []models.ArticleVotes `json:"articles"`
	Err      string                `json:"err,omitempty"`
}
