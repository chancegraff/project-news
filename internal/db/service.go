package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL adapter
)

// Service holds the database store
type Service struct {
	Store    *Store
	Articles *Articles
}

// NewService will instantiate a Service
func NewService() (*Service, error) {
	// Create store
	store, err := NewStore()
	if err != nil {
		return nil, err
	}

	// Create service
	service := Service{Store: store}

	// Bind children
	service.Articles = NewArticles(&service)

	return &service, nil
}

// Stop will close connections and stop the service
func (s *Service) Stop() {
	s.Store.Close()
}
