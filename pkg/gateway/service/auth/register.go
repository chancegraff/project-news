package auth

import (
	pba "github.com/chancegraff/project-news/api/auth"
)

// Register ...
func (s *service) Register(email string, password string) (*pba.User, error) {
	return s.Proxy.Auth.Register(email, password)
}
