package interfaces

import "github.com/chancegraff/project-news/internal/models"

// GetRequest ...
type GetRequest struct {
	ID int `json:"id"`
}

// GetResponse ...
type GetResponse struct {
	Article models.Article `json:"article"`
	Err     string         `json:"err,omitempty"`
}
