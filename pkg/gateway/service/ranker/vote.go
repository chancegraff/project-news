package ranker

import (
	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Vote ...
func (s *service) Vote(articleID, userID string) (*pbr.ArticleVotes, error) {
	return s.Proxy.Ranker.Vote(articleID, userID)
}
