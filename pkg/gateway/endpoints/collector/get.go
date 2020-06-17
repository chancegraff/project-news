package collector

import (
	"context"
	"errors"

	pbc "github.com/chancegraff/project-news/api/collector"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeGetEndpoint ...
func MakeGetEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbc.GetRequest)
		article, err := svc.Collector.Get(int(req.Id))
		if err != nil {
			return pbc.GetResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return pbc.GetResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

// Get ...
func (e Endpoints) Get(ctx context.Context, id int) (*pbc.Article, error) {
	req := &pbc.GetRequest{Id: int32(id)}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(*pbc.GetResponse)
	if getResp.Err != "" {
		return nil, errors.New(getResp.Err)
	}
	return getResp.Article, nil
}
