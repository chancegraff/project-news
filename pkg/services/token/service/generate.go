package service

import (
	"github.com/chancegraff/project-news/internal/encoding"
	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/internal/utils"
)

// Generate ...
func (s *service) Generate(identifiers models.Identifiers, client models.Client) (string, error) {
	// Generate hash
	hash, err := encoding.Hash(identifiers, client.IP)
	if err != nil {
		return "", err
	}

	// Update client
	client.Hash = hash
	client.ExpiredAt = utils.Tomorrow()

	// Save client
	client, err = s.Manager.FirstOrCreate(client)
	if err != nil {
		return "", err
	}

	// Return final hash
	return hash, nil
}
