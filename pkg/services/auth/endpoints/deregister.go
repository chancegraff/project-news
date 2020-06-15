package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	"github.com/chancegraff/project-news/pkg/services/auth/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeDeregisterEndpoint ...
func MakeDeregisterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.DeregisterRequest)
		user, err := svc.Deregister(req.UserID)
		if err != nil {
			return transports.DeregisterResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return transports.DeregisterResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Deregister ...
func (e Endpoints) Deregister(ctx context.Context, userID string) (models.User, error) {
	req := transports.DeregisterRequest{UserID: userID}
	resp, err := e.DeregisterEndpoint(ctx, req)
	if err != nil {
		return models.User{}, err
	}
	deregisterResp := resp.(transports.DeregisterResponse)
	if deregisterResp.Err != "" {
		return models.User{}, errors.New(deregisterResp.Err)
	}
	return deregisterResp.User, nil
}
