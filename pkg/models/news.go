package models

import (
	"time"
)

// NewsArticle ...
type NewsArticle struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Thumbnail   string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

// NewsResponse ...
type NewsResponse struct {
	Status  string        `json:"status"`
	Results []NewsArticle `json:"articles"`
}
