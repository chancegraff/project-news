package token

import (
	"context"
	"errors"

	pbt "github.com/chancegraff/project-news/api/token"
)

// Verify ...
func (p *proxy) Verify(identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	verifyResp, err := p.Client.Verify(context.Background(), &pbt.VerifyRequest{Identifiers: identifiers, Client: client})
	if err != nil {
		return "", err
	}
	if err := verifyResp.Err; err != "" {
		return "", errors.New(verifyResp.Err)
	}
	return verifyResp.Hash, nil
}
