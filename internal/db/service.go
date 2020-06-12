package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL adapter
)

// Service holds the database store
type Service struct {
	Store    *Store
	Articles *Articles
}

// Start will open connections and start the service
func (s *Service) Start() error {
	// Create store
	store, err := NewStore()
	if err != nil {
		return err
	}

	// Bind children
	s.Store = store
	s.Articles = &Articles{s}

	return nil
}

// Stop will close connections and stop the service
func (s *Service) Stop() {
	s.Store.Stop()
}

// NewService will instantiate a database Service and start it
func NewService() (*Service, error) {
	service := Service{}
	err := service.Start()
	return &service, err
}
