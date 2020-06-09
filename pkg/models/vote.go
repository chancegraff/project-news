package models

import "github.com/jinzhu/gorm"

// Vote ...
type Vote struct {
	UserID    string `json:"user"`
	ArticleID string `json:"article"`
	gorm.Model
}
