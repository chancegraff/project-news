package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article ...
type Article struct {
	Title       string    `json:"Title"`
	URL         string    `json:"URL"`
	Thumbnail   string    `json:"Thumbnail"`
	PublishedAt time.Time `json:"PublishedAt"`
	gorm.Model
}

// ArticleResponse ...
type ArticleResponse struct {
	Status  string    `json:"status"`
	Results []Article `json:"results"`
}
