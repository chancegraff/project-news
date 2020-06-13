package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	"github.com/go-kit/kit/endpoint"
)

func makeGetEndpoint(svc service.Service) endpoint.Endpoint {
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
