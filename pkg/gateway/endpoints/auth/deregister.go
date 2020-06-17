package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeDeregisterEndpoint ...
func MakeDeregisterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pba.DeregisterRequest)
		user, err := svc.Auth.Deregister(req.UserID)
		if err != nil {
			return pba.DeregisterResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return pba.DeregisterResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Deregister ...
func (e Endpoints) Deregister(ctx context.Context, userID string) (*pba.User, error) {
	req := &pba.DeregisterRequest{UserID: userID}
	resp, err := e.DeregisterEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	deregisterResp := resp.(*pba.DeregisterResponse)
	if deregisterResp.Err != "" {
		return nil, errors.New(deregisterResp.Err)
	}
	return deregisterResp.User, nil
}
