package service

import "github.com/chancegraff/project-news/internal/models"

// Get will find and return a single article from the database that matches the ID
func (s *service) Get(id int) (models.Article, error) {
	return s.Manager.First(id)
}
