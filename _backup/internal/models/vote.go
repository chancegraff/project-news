package models

// Vote is a representation of a vote in the database
type Vote struct {
	UserID    string `json:"user"`
	ArticleID string `json:"article"`
	Base
}
