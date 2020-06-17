package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeArticlesEndpoint ...
func MakeArticlesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbr.ArticlesRequest)
		articles, err := svc.Ranker.Articles(req.ArticleIDs)
		if err != nil {
			return pbr.ArticlesResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return pbr.ArticlesResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

// Articles ...
func (e Endpoints) Articles(ctx context.Context, articleIDs []string) ([]*pbr.ArticleVotes, error) {
	req := &pbr.ArticlesRequest{ArticleIDs: articleIDs}
	resp, err := e.ArticlesEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	articlesResp := resp.(*pbr.ArticlesResponse)
	if articlesResp.Err != "" {
		return nil, errors.New(articlesResp.Err)
	}
	return articlesResp.Articles, nil
}
