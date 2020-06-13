package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/ranker/interfaces"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/go-kit/kit/endpoint"
)

func makeArticlesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(interfaces.ArticlesRequest)
		articles, err := svc.Articles(req.ArticleIDs)
		if err != nil {
			return interfaces.ArticlesResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return interfaces.ArticlesResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}
