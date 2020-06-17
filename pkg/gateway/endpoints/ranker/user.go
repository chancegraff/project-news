package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeUserEndpoint ...
func MakeUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbr.UserRequest)
		user, err := svc.Ranker.User(req.UserID)
		if err != nil {
			return pbr.UserResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return pbr.UserResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// User ...
func (e Endpoints) User(ctx context.Context, userID string) (*pbr.UserVotes, error) {
	req := &pbr.UserRequest{UserID: userID}
	resp, err := e.UserEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	userResp := resp.(*pbr.UserResponse)
	if userResp.Err != "" {
		return nil, errors.New(userResp.Err)
	}
	return userResp.User, nil
}
