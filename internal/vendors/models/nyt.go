package models

import (
	"encoding/json"
	"time"
)

type multimedia struct {
	URL    string `json:"url"`
	Format string `json:"format"`
}

// NewYorkTimesArticle ...
type NewYorkTimesArticle struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt time.Time `json:"published_date"`
}

// UnmarshalJSON ...
func (art *NewYorkTimesArticle) UnmarshalJSON(buf []byte) error {
	type Alias NewYorkTimesArticle
	raw := &struct {
		Multimedia []multimedia `json:"multimedia"`
		*Alias
	}{
		Alias: (*Alias)(art),
	}
	if err := json.Unmarshal(buf, &raw); err != nil {
		return err
	}

	for _, m := range raw.Multimedia {
		if m.Format == "thumbLarge" {
			art.Thumbnail = m.URL
		}
	}

	return nil
}

// NewYorkTimesResponse ...
type NewYorkTimesResponse struct {
	Status  string                `json:"status"`
	Results []NewYorkTimesArticle `json:"results"`
}
