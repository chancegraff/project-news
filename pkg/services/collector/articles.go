package collector

import (
	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/db/middleware"
	"github.com/chancegraff/project-news/internal/models"
)

// Articles interfaces with the database for the articles table
type Articles struct {
	Store *db.Store
}

// List will return an array of articles sorted by published_at
func (a *Articles) List(offset, limit int) ([]models.Article, error) {
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

// First will find and return the first article in the database that matches the ID
func (a *Articles) First(id int) (models.Article, error) {
	// Get article
	var article models.Article
	err := a.Store.Database.First(&article, id).Error
	if err != nil {
		return article, err
	}

	// Return result
	return article, nil
}

// FirstOrCreate will create a new record or find an existing record and return either
func (a *Articles) FirstOrCreate(article *models.Article) models.Article {
	buffer := *article
	a.Store.Database.Where(models.Article{URL: article.URL}).FirstOrCreate(buffer)
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
