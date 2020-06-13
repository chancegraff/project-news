package manager

import "github.com/chancegraff/project-news/internal/models"

// First will find and return the first article in the database that matches the ID
func (a *Manager) First(id int) (models.Article, error) {
	// Get article
	var article models.Article
	err := a.Store.Database.First(&article, id).Error
	if err != nil {
		return article, err
	}

	// Return result
	return article, nil
}
