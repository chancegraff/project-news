package collector

import (
	"context"
	"errors"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// All ...
func (p proxy) All(ctx context.Context, offset int) ([]*pbc.Article, error) {
	allResp, err := p.Client.All(ctx, &pbc.AllRequest{Offset: int32(offset)})
	if err != nil {
		return nil, err
	}
	if err := allResp.Err; err != "" {
		return nil, errors.New(allResp.Err)
	}
	articles := allResp.Articles
	return articles, nil
}
