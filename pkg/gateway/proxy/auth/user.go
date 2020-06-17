package auth

import (
	"context"
	"errors"

	pba "github.com/chancegraff/project-news/api/auth"
)

// User ...
func (p proxy) User(userID string) (*pba.User, error) {
	userResp, err := p.Client.User(context.Background(), &pba.UserRequest{UserID: userID})
	if err != nil {
		return nil, err
	}
	if err := userResp.Err; err != "" {
		return nil, errors.New(userResp.Err)
	}
	return userResp.User, nil
}
