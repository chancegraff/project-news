package auth

import (
	pba "github.com/chancegraff/project-news/api/auth"
)

// Deregister ...
func (s *service) Deregister(userID string) (*pba.User, error) {
	return s.Proxy.Auth.Deregister(userID)
}
