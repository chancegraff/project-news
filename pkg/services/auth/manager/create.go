package manager

import "github.com/chancegraff/project-news/internal/models"

// Create will take an email string and hash byte array and save them to a new record
func (m *Manager) Create(email string, hash []byte) (models.User, error) {
	// Instantiate instance
	user := models.User{
		Credentials: models.Credentials{
			Email:    email,
			Password: string(hash),
		},
	}

	// Save instance to database
	err := m.Store.Database.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
