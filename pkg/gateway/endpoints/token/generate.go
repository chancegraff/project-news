package token

import (
	"context"
	"errors"

	pbt "github.com/chancegraff/project-news/api/token"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeGenerateEndpoint ...
func MakeGenerateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbt.GenerateRequest)
		hash, err := svc.Token.Generate(req.Identifiers, req.Client)
		if err != nil {
			return pbt.GenerateResponse{
				Hash: hash,
				Err:  err.Error(),
			}, nil
		}
		return pbt.GenerateResponse{
			Hash: hash,
			Err:  "",
		}, nil
	}
}

// Generate ...
func (e Endpoints) Generate(ctx context.Context, identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	req := &pbt.GenerateRequest{Identifiers: identifiers, Client: client}
	resp, err := e.GenerateEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	generateResp := resp.(*pbt.GenerateResponse)
	if generateResp.Err != "" {
		return "", errors.New(generateResp.Err)
	}
	return generateResp.Hash, nil
}
