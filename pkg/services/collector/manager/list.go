package manager

import (
	"github.com/chancegraff/project-news/internal/db/middleware"
	"github.com/chancegraff/project-news/internal/models"
)

// List will return an array of articles sorted by published_at
func (a *Manager) List(offset, limit int) ([]models.Article, error) {
	// Setup store
	database := middleware.WithLimit(
		middleware.WithOffset(
			a.Store.Database,
			offset,
		),
		limit,
	)

	// Retrieve articles
	var articles []models.Article
	err := database.Order("published_at desc").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}
