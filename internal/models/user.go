package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite adapter
)

// Credentials is an email and password combination
type Credentials struct {
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password" gorm:"not null"`
}

// User is the representation of a user in the database
type User struct {
	Base
	Credentials
	VerifiedAt time.Time `json:"verified_at"`
}
