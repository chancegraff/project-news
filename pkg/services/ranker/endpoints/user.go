package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
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

// User ...
func (e Endpoints) User(ctx context.Context, userID string) (models.UserVotes, error) {
	req := transports.UserRequest{UserID: userID}
	resp, err := e.UserEndpoint(ctx, req)
	if err != nil {
		return models.UserVotes{}, err
	}
	userResponse := resp.(transports.UserResponse)
	if userResponse.Err != "" {
		return userResponse.User, errors.New(userResponse.Err)
	}
	return userResponse.User, nil
}
