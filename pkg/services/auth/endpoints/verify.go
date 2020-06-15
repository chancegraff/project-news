package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	"github.com/chancegraff/project-news/pkg/services/auth/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.VerifyRequest)
		user, err := svc.Verify(req.Email, req.Password)
		if err != nil {
			return transports.VerifyResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return transports.VerifyResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Verify ...
func (e Endpoints) Verify(ctx context.Context, email string, password string) (models.User, error) {
	req := transports.VerifyRequest{Email: email, Password: password}
	resp, err := e.VerifyEndpoint(ctx, req)
	if err != nil {
		return models.User{}, err
	}
	verifyResp := resp.(transports.VerifyResponse)
	if verifyResp.Err != "" {
		return models.User{}, errors.New(verifyResp.Err)
	}
	return verifyResp.User, nil
}
