package collector

import (
	"context"
	"errors"

	pbc "github.com/chancegraff/project-news/api/collector"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeAllEndpoint ...
func MakeAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbc.AllRequest)
		articles, err := svc.Collector.All(int(req.Offset))
		if err != nil {
			return pbc.AllResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return pbc.AllResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

// All ...
func (e Endpoints) All(ctx context.Context, offset int) ([]*pbc.Article, error) {
	req := &pbc.AllRequest{Offset: int32(offset)}
	resp, err := e.AllEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	allResp := resp.(*pbc.AllResponse)
	if allResp.Err != "" {
		return nil, errors.New(allResp.Err)
	}
	return allResp.Articles, nil
}
