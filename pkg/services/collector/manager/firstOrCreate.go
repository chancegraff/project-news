package manager

import (
	"github.com/chancegraff/project-news/internal/models"
)

// FirstOrCreate will create a new record or find an existing record and return either
func (a *Manager) FirstOrCreate(article models.Article) (models.Article, error) {
	buffer := article
	err := a.Store.Database.Where(models.Article{URL: article.URL}).FirstOrCreate(&buffer).Error
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
