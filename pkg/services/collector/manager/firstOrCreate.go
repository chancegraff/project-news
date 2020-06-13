package manager

import "github.com/chancegraff/project-news/internal/models"

// FirstOrCreate will create a new record or find an existing record and return either
func (a *Manager) FirstOrCreate(article *models.Article) models.Article {
	buffer := *article
	a.Store.Database.Where(models.Article{URL: article.URL}).FirstOrCreate(buffer)
	return buffer
}
