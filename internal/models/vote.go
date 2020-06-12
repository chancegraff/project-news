package models

// Vote is a representation of a vote in the database
type Vote struct {
	UserID    string `json:"user"`
	ArticleID string `json:"article"`
	Base
}

// UserVotes is a representation of the votes a single user has in the database
type UserVotes struct {
	UserID string `json:"user"`
	Votes  []Vote `json:"votes"`
}

// ArticleVotes is a representation of the votes a single article has in the database
type ArticleVotes struct {
	ArticleID string `json:"article"`
	Votes     int    `json:"votes"`
}
