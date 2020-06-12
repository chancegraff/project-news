package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite adapter
)

// Credentials ...
type Credentials struct {
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password" gorm:"not null"`
}

// User ...
type User struct {
	Base
	Credentials
}

// UserWithArticles ...
type UserWithArticles struct {
	ID       string
	Email    string
	Articles []string
}
