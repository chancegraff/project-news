package token

import (
	"context"
	"errors"

	pbt "github.com/chancegraff/project-news/api/token"
)

// Generate ...
func (p proxy) Generate(identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	generateResp, err := p.Client.Generate(context.Background(), &pbt.GenerateRequest{Identifiers: identifiers, Client: client})
	if err != nil {
		return "", err
	}
	if err := generateResp.Err; err != "" {
		return "", errors.New(generateResp.Err)
	}
	return generateResp.Hash, nil
}
