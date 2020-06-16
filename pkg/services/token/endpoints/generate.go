package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/pkg/services/token/service"
	"github.com/chancegraff/project-news/pkg/services/token/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeGenerateEndpoint ...
func MakeGenerateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.GenerateRequest)
		hash, err := svc.Generate(req.Identifiers, req.Client)
		if err != nil {
			return transports.GenerateResponse{
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

// Generate ...
func (e Endpoints) Generate(ctx context.Context, identifiers models.Identifiers, client models.Client) (string, error) {
	req := transports.GenerateRequest{Identifiers: identifiers, Client: client}
	resp, err := e.GenerateEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	generateResp := resp.(transports.GenerateResponse)
	if generateResp.Err != "" {
		return generateResp.Hash, errors.New(generateResp.Err)
	}
	return generateResp.Hash, nil
}
