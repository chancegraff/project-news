package vendors

import (
	"encoding/json"
	"os"

	"github.com/chancegraff/project-news/pkg/models"
)

var newsURL = "https://newsapi.org/v2/everything?q=politics&domains=msnbc.com,cnn.com,washingtonpost.com&sortBy=publishedAt&language=en&apiKey="

// News returns data from News API
func News() (*models.NewsResponse, error) {
	newsToken := os.Getenv("API_TOKEN_NEWS")
	body, err := getArticles(newsURL, newsToken)
	if err != nil {
		return nil, err
	}

	var nRes models.NewsResponse
	err = json.Unmarshal(*body, &nRes)
	if err != nil {
		return nil, err
	}

	return &nRes, nil
}
