package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeArticlesEndpoint ...
func MakeArticlesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.ArticlesRequest)
		articles, err := svc.Articles(req.ArticleIDs)
		if err != nil {
			return transports.ArticlesResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return transports.ArticlesResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

// Articles ...
func (e Endpoints) Articles(ctx context.Context, articleIDs []string) ([]models.ArticleVotes, error) {
	req := transports.ArticlesRequest{ArticleIDs: articleIDs}
	resp, err := e.ArticlesEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	articlesResp := resp.(transports.ArticlesResponse)
	if articlesResp.Err != "" {
		return nil, errors.New(articlesResp.Err)
	}
	return articlesResp.Articles, nil
}
