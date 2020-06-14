package manager

import (
	"github.com/chancegraff/project-news/internal/db"
)

// Manager interfaces with the database for the articles table
type Manager struct {
	Store *db.Store
}

// NewManager creates a database store and returns a manager for it
func NewManager() Manager {
	store, err := db.NewStore()
	if err != nil {
		panic(err)
	}
	return Manager{
		Store: store,
	}
}
