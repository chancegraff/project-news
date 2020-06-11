package models

// Vote ...
type Vote struct {
	UserID    string `json:"user"`
	ArticleID string `json:"article"`
	Base
}
