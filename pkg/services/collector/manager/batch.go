package manager

import "github.com/chancegraff/project-news/internal/models"

// Batch will create records from an array and return them
func (a *Manager) Batch(articles *[]models.Article) []models.Article {
	buffer, result := *articles, make([]models.Article, len(*articles))
	for _, article := range buffer {
		result = append(result, a.FirstOrCreate(&article))
	}
	return result
}
