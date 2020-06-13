package manager

import (
	"github.com/chancegraff/project-news/internal/db"
)

// Manager interfaces with the database for the articles table
type Manager struct {
	Store *db.Store
}
