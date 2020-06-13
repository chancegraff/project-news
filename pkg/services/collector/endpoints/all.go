package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/collector/interfaces"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/go-kit/kit/endpoint"
)

func makeAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(interfaces.AllRequest)
		articles, err := svc.All(req.Offset)
		if err != nil {
			return interfaces.AllResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return interfaces.AllResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}
