package service

import (
	"errors"

	"github.com/chancegraff/project-news/internal/encoding"
	"github.com/chancegraff/project-news/internal/models"
)

// Verify ...
func (s *service) Verify(identifiers models.Identifiers, client models.Client) (string, error) {
	// Retrieve client
	client, err := s.Manager.First(client)
	if err != nil {
		return "", err
	}

	// Generate hash
	hash, err := encoding.Hash(identifiers, client.IP)
	if err != nil {
		return "", err
	}

	// Verify hashes match
	if hash != client.Hash {
		return "", errors.New("invalid hash")
	}

	// Return generated hash
	return hash, nil
}
