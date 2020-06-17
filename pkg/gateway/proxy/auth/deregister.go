package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"
)

// Deregister ...
func (p proxy) Deregister(userID string) (*pba.User, error) {
	deregisterResp, err := p.Client.Deregister(context.Background(), &pba.DeregisterRequest{UserID: userID})
	if err != nil {
		return nil, err
	}
	if err := deregisterResp.Err; err != "" {
		return nil, errors.New(deregisterResp.Err)
	}
	return deregisterResp.User, nil
}
