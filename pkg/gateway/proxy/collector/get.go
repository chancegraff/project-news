package collector

import (
	"context"
	"errors"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// Get ...
func (p *proxy) Get(id int) (*pbc.Article, error) {
	getResp, err := p.Client.Get(context.Background(), &pbc.GetRequest{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	if err := getResp.Err; err != "" {
		return nil, errors.New(getResp.Err)
	}
	article := getResp.Article
	return article, nil
}
