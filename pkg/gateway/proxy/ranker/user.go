package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// User ...
func (p proxy) User(userID string) (*pbr.UserVotes, error) {
	userResp, err := p.Client.User(context.Background(), &pbr.UserRequest{UserID: userID})
	if err != nil {
		return nil, err
	}
	if err := userResp.Err; err != "" {
		return nil, errors.New(userResp.Err)
	}
	return userResp.User, nil
}
