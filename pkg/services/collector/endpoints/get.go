package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeGetEndpoint ...
func MakeGetEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.GetRequest)
		article, err := svc.Get(req.ID)
		if err != nil {
			return transports.GetResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return transports.GetResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

// Get ...
func (e Endpoints) Get(ctx context.Context, id int) (models.Article, error) {
	req := transports.GetRequest{ID: id}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return models.Article{}, err
	}
	getResp := resp.(transports.GetResponse)
	if getResp.Err != "" {
		return models.Article{}, errors.New(getResp.Err)
	}
	return getResp.Article, nil
}
