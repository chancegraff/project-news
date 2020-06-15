package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	"github.com/chancegraff/project-news/pkg/services/auth/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeRegisterEndpoint ...
func MakeRegisterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.RegisterRequest)
		user, err := svc.Register(req.Email, req.Password)
		if err != nil {
			return transports.RegisterResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return transports.RegisterResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Register ...
func (e Endpoints) Register(ctx context.Context, email string, password string) (models.User, error) {
	req := transports.RegisterRequest{Email: email, Password: password}
	resp, err := e.RegisterEndpoint(ctx, req)
	if err != nil {
		return models.User{}, err
	}
	registerResp := resp.(transports.RegisterResponse)
	if registerResp.Err != "" {
		return models.User{}, errors.New(registerResp.Err)
	}
	return registerResp.User, nil
}
