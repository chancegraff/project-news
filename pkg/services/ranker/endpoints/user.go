package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeUserEndpoint ...
func MakeUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.UserRequest)
		user, err := svc.User(req.UserID)
		if err != nil {
			return transports.UserResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return transports.UserResponse{
			User: user,
			Err:  "",
		}, nil
	}
}
