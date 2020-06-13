package interfaces

import "github.com/chancegraff/project-news/internal/models"

// AllRequest ...
type AllRequest struct {
	Offset int `json:"offset"`
}

// AllResponse ...
type AllResponse struct {
	Articles []models.Article `json:"articles"`
	Err      string           `json:"err,omitempty"`
}
