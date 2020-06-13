package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/ranker/interfaces"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/go-kit/kit/endpoint"
)

func makeUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(interfaces.UserRequest)
		user, err := svc.User(req.UserID)
		if err != nil {
			return interfaces.UserResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return interfaces.UserResponse{
			User: user,
			Err:  "",
		}, nil
	}
}
