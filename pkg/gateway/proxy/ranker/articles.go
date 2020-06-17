package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Articles ...
func (p proxy) Articles(articleIDs []string) ([]*pbr.ArticleVotes, error) {
	articlesResp, err := p.Client.Articles(context.Background(), &pbr.ArticlesRequest{ArticleIDs: articleIDs})
	if err != nil {
		return nil, err
	}
	if err := articlesResp.Err; err != "" {
		return nil, errors.New(articlesResp.Err)
	}
	return articlesResp.Articles, nil
}
