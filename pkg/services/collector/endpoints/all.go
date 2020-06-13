package endpoints

import (
	"context"
	"log"

	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	"github.com/go-kit/kit/endpoint"
)

func makeAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		log.Println("all endpoint")
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
