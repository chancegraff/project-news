package token

import (
	"context"
	"errors"

	pbt "github.com/chancegraff/project-news/api/token"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeVerifyEndpoint ...
func MakeVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbt.VerifyRequest)
		hash, err := svc.Token.Verify(req.Identifiers, req.Client)
		if err != nil {
			return pbt.VerifyResponse{
				Hash: hash,
				Err:  err.Error(),
			}, nil
		}
		return pbt.VerifyResponse{
			Hash: hash,
			Err:  "",
		}, nil
	}
}

// Verify ...
func (e Endpoints) Verify(ctx context.Context, identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	req := &pbt.VerifyRequest{Identifiers: identifiers, Client: client}
	resp, err := e.VerifyEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	verifyResp := resp.(*pbt.VerifyResponse)
	if verifyResp.Err != "" {
		return "", errors.New(verifyResp.Err)
	}
	return verifyResp.Hash, nil
}
