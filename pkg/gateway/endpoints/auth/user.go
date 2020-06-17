package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeUserEndpoint ...
func MakeUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pba.UserRequest)
		user, err := svc.Auth.User(req.UserID)
		if err != nil {
			return pba.UserResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return pba.UserResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// User ...
func (e Endpoints) User(ctx context.Context, userID string) (*pba.User, error) {
	req := &pba.UserRequest{UserID: userID}
	resp, err := e.UserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	userResp := resp.(*pba.UserResponse)
	if userResp.Err != "" {
		return nil, errors.New(userResp.Err)
	}
	return userResp.User, nil
}
