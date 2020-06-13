package service

import (
	"log"
	"time"

	"github.com/chancegraff/project-news/internal/vendors"
)

// Collect ...
func (s *service) Collect() {
	articles := vendors.Get()
	s.manager.Batch(articles)
	log.Println("Server updated articles")
	time.Sleep(5 * time.Minute)
	go s.Collect()
}
