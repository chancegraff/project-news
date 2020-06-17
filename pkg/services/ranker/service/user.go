package service

import "github.com/chancegraff/project-news/internal/models"

// User will take a user ID and returns an array of article IDs associated with it
func (s *service) User(userID string) (models.UserVotes, error) {
	return s.Manager.GetVotesForUser(userID)
}
