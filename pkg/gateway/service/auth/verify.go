package auth

import (
	pba "github.com/chancegraff/project-news/api/auth"
)

// Verify ...
func (s *service) Verify(email string, password string) (*pba.User, error) {
	return s.Proxy.Auth.Verify(email, password)
}
