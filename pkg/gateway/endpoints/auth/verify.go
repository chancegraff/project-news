package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pba.VerifyRequest)
		user, err := svc.Auth.Verify(req.Email, req.Password)
		if err != nil {
			return pba.VerifyResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return pba.VerifyResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Verify ...
func (e Endpoints) Verify(ctx context.Context, email string, password string) (*pba.User, error) {
	req := &pba.VerifyRequest{Email: email, Password: password}
	resp, err := e.VerifyEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	verifyResp := resp.(*pba.VerifyResponse)
	if verifyResp.Err != "" {
		return nil, errors.New(verifyResp.Err)
	}
	return verifyResp.User, nil
}
