package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"
)

// Verify ...
func (p *proxy) Verify(email string, password string) (*pba.User, error) {
	verifyResp, err := p.Client.Verify(context.Background(), &pba.VerifyRequest{Email: email, Password: password})
	if err != nil {
		return nil, err
	}
	if err := verifyResp.Err; err != "" {
		return nil, errors.New(verifyResp.Err)
	}
	return verifyResp.User, nil
}
