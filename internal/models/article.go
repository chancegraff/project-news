package models

import (
	"time"
)

// Article is a representation of an article in the database
type Article struct {
	Title       string    `json:"Title"`
	URL         string    `json:"URL"`
	Thumbnail   string    `json:"Thumbnail"`
	PublishedAt time.Time `json:"PublishedAt"`
	Base
}
