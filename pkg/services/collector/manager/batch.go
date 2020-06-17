package manager

import "github.com/chancegraff/project-news/internal/models"

// Batch will create records from an array and return them
func (a *Manager) Batch(articles []models.Article) []models.Article {
	buffer, result := articles, make([]models.Article, len(articles))
	for _, article := range buffer {
		a, err := a.FirstOrCreate(article)
		if err != nil {
			break
		}
		result = append(result, a)
	}
	return result
}
