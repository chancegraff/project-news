package endpoints

import (
	"context"

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
