package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeRegisterEndpoint ...
func MakeRegisterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pba.RegisterRequest)
		user, err := svc.Auth.Register(req.Email, req.Password)
		if err != nil {
			return pba.RegisterResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return pba.RegisterResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Register ...
func (e Endpoints) Register(ctx context.Context, email string, password string) (*pba.User, error) {
	req := &pba.RegisterRequest{Email: email, Password: password}
	resp, err := e.RegisterEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	registerResp := resp.(*pba.RegisterResponse)
	if registerResp.Err != "" {
		return nil, errors.New(registerResp.Err)
	}
	return registerResp.User, nil
}
