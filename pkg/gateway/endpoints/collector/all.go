package collector

import (
	"context"
	"errors"
	"log"

	pbc "github.com/chancegraff/project-news/api/collector"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeAllEndpoint ...
func MakeAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		log.Println("All endpoint 0")
		req := request.(*pbc.AllRequest)
		log.Println("All endpoint 1")
		articles, err := svc.Collector.All(int(req.Offset))
		log.Println("All endpoint 2")
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
	log.Println("All 1")
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
