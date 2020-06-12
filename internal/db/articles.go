package db

import (
	"github.com/chancegraff/project-news/internal/db/middleware"
	"github.com/chancegraff/project-news/internal/models"
)

// Articles interfaces with the database for the articles table
type Articles struct {
	service *Service
}

// List will return an array of articles sorted by published_at
func (a *Articles) List(offset, limit int) ([]models.Article, error) {
	// Setup store
	store := middleware.WithLimit(
		middleware.WithOffset(
			a.service.Store.Database,
			offset,
		),
		limit,
	)

	// Retrieve articles
	var articles []models.Article
	err := store.Order("published_at desc").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// FirstOrCreate will create a new record or find an existing record and return either
func (a *Articles) FirstOrCreate(article *models.Article) models.Article {
	buffer := *article
	a.service.Store.Database.Where(models.Article{URL: article.URL}).FirstOrCreate(buffer)
	return buffer
}

// Batch will create records from an array and return them
func (a *Articles) Batch(articles *[]models.Article) []models.Article {
	buffer, result := *articles, make([]models.Article, len(*articles))
	for _, article := range buffer {
		result = append(result, a.FirstOrCreate(&article))
	}
	return result
}

// NewArticles will instantiate a new articles service
func NewArticles(s *Service) *Articles {
	return &Articles{s}
}
