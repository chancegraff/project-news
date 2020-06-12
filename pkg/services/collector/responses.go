package collector

import "github.com/chancegraff/project-news/internal/models"

type getResponse struct {
	Article models.Article `json:"article"`
	Err     string         `json:"err,omitempty"`
}

type allResponse struct {
	Articles []models.Article `json:"articles"`
	Err      string           `json:"err,omitempty"`
}

// Proxies

type articlesResponse struct {
	Articles []models.ArticleVotes `json:"articles"`
	Err      string                `json:"err,omitempty"`
}
