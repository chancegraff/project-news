package ranker

import (
	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Articles ...
func (s *service) Articles(articleIDs []string) ([]*pbr.ArticleVotes, error) {
	return s.Proxy.Ranker.Articles(articleIDs)
}
