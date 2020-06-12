package vendors

import (
	"encoding/json"
	"os"

	"github.com/chancegraff/project-news/pkg/models"
)

var newYorkTimesURL = "https://api.nytimes.com/svc/topstories/v2/politics.json?api-key="

// NewYorkTimes returns data from NYTimes API
func NewYorkTimes() (*models.NewYorkTimesResponse, error) {
	nytToken := os.Getenv("API_TOKEN_NYTIMES")
	body, err := getArticles(newYorkTimesURL, nytToken)
	if err != nil {
		return nil, err
	}

	var nRes models.NewYorkTimesResponse
	err = json.Unmarshal(*body, &nRes)
	if err != nil {
		return nil, err
	}

	return &nRes, nil
}
