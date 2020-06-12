package models

import "time"

// Base can be used as a base for GORM models
type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
