package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/token/service"
	"github.com/chancegraff/project-news/pkg/services/token/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.VerifyRequest)
		hash, err := svc.Verify(req.Identifiers, req.Client)
		if err != nil {
			return transports.VerifyResponse{
				Hash: hash,
				Err:  err.Error(),
			}, nil
		}
		return transports.GenerateResponse{
			Hash: hash,
			Err:  "",
		}, nil
	}
}

// Verify ...
func (e Endpoints) Verify(ctx context.Context, identifiers models.Identifiers, client models.Client) (string, error) {
	req := transports.VerifyRequest{Identifiers: identifiers, Client: client}
	resp, err := e.GenerateEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	verifyResp := resp.(transports.VerifyResponse)
	if verifyResp.Err != "" {
		return verifyResp.Hash, errors.New(verifyResp.Err)
	}
	return verifyResp.Hash, nil
}
