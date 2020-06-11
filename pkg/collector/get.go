package collector

import (
	"encoding/json"

	"github.com/chancegraff/project-news/pkg/models"
)

type bodyGet struct {
	ID string
}

func get(data interface{}) ([]byte, error) {
	// Typecast into body
	bd := data.(bodyGet)

	// Get article
	var article models.Article
	err := store.First(&article, bd.ID).Error
	if err != nil {
		return nil, err
	}

	// Return data
	return json.Marshal(article)
}
