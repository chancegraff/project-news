package ranker

import "github.com/chancegraff/project-news/internal/models"

type articlesResponse struct {
	Articles []models.ArticleVotes `json:"articles"`
	Err      string                `json:"err,omitempty"`
}

type userResponse struct {
	User models.UserVotes `json:"user"`
	Err  string           `json:"err,omitempty"`
}

type voteResponse struct {
	Article models.ArticleVotes `json:"article"`
	Err     string              `json:"err,omitempty"`
}
