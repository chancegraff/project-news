package auth

import (
	pba "github.com/chancegraff/project-news/api/auth"
)

// User ...
func (s *service) User(userID string) (*pba.User, error) {
	return s.Proxy.Auth.User(userID)
}
