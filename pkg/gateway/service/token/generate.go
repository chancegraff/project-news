package token

import (
	pbt "github.com/chancegraff/project-news/api/token"
)

// Generate ...
func (s *service) Generate(identifiers *pbt.Identifiers, client *pbt.Client) (string, error) {
	return s.Proxy.Token.Generate(identifiers, client)
}
