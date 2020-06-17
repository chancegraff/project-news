package models

import "time"

// Base is the attributes needed to turn a model into a table in the database
type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
