package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"
)

// Register ...
func (p *proxy) Register(email string, password string) (*pba.User, error) {
	registerResp, err := p.Client.Register(context.Background(), &pba.RegisterRequest{Email: email, Password: password})
	if err != nil {
		return nil, err
	}
	if err := registerResp.Err; err != "" {
		return nil, errors.New(registerResp.Err)
	}
	return registerResp.User, nil
}
