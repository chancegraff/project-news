package collector

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		article, err := svc.Get(req.ID)
		if err != nil {
			return getResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return getResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

func makeAllEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(allRequest)
		articles, err := svc.All(req.Offset)
		if err != nil {
			return allResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return allResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}
