package interfaces

import "github.com/chancegraff/project-news/internal/models"

// UserRequest ...
type UserRequest struct {
	UserID string `json:"user"`
}

// UserResponse ...
type UserResponse struct {
	User models.UserVotes `json:"user"`
	Err  string           `json:"err,omitempty"`
}
