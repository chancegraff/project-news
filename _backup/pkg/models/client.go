package models

import (
	"time"

	"github.com/google/uuid"
)

// Client represents a unique individual
type Client struct {
	Hash      string    `json:"hash" gorm:"unique_index"`
	UserID    string    `json:"user"`
	Secret    uuid.UUID `json:"secret"`
	ExpiredAt time.Time `json:"expiredAt"`
	Base
}

// Identifiers are used to construct a client hash
type Identifiers struct {
	Software string `json:"software"` //navigator.platform
	Browser  string `json:"browser"`  //navigator.userAgent
	Language string `json:"language"` //navigator.language
	Width    string `json:"width"`    //screen.width
	Height   string `json:"height"`   //screen.height
	Colors   string `json:"colors"`   //screen.colorDepth
	Pixels   string `json:"pixels"`   //screen.pixelDepth
}
