package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeAllEndpoint ...
func MakeAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.AllRequest)
		articles, err := svc.All(req.Offset)
		if err != nil {
			return transports.AllResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return transports.AllResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

// All ...
func (e Endpoints) All(ctx context.Context, offset int) ([]models.Article, error) {
	req := transports.AllRequest{Offset: offset}
	resp, err := e.AllEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	allResp := resp.(transports.AllResponse)
	if allResp.Err != "" {
		return nil, errors.New(allResp.Err)
	}
	return allResp.Articles, nil
}
