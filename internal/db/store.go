package db

import (
	"os"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL adapter
)

// Store is the connection to the database
type Store struct {
	Database *gorm.DB
}

// Start will instantiate a database connection
func (s *Store) Start() error {
	// Open connection
	database, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	// Set connection
	s.Database = database

	return nil
}

// Migrate will run automigrate for models
func (s *Store) Migrate() {
	s.Database.AutoMigrate(&models.Article{})
	s.Database.AutoMigrate(&models.Client{})
	s.Database.AutoMigrate(&models.User{})
	s.Database.AutoMigrate(&models.Vote{})
}

// Close will close the connection to the database
func (s *Store) Close() {
	s.Database.Close()
}

// NewStore will instantiate a Store
func NewStore() (*Store, error) {
	store := Store{}

	// Start connection
	err := store.Start()
	if err != nil {
		return nil, err
	}

	// Run migrations
	store.Migrate()

	return &store, nil
}
