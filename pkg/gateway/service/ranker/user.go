package ranker

import (
	pbr "github.com/chancegraff/project-news/api/ranker"
)

// User ...
func (s *service) User(userID string) (*pbr.UserVotes, error) {
	return s.Proxy.Ranker.User(userID)
}
