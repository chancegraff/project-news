package token

import (
	pbt "github.com/chancegraff/project-news/api/token"
)

// Verify ...
func (s *service) Verify(identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	return s.Proxy.Token.Verify(identifiers, client)
}
