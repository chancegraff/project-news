package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/collector/interfaces"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/go-kit/kit/endpoint"
)

func makeGetEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(interfaces.GetRequest)
		article, err := svc.Get(req.ID)
		if err != nil {
			return interfaces.GetResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return interfaces.GetResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}
