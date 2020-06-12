package vendors

import (
	"github.com/chancegraff/project-news/internal/models"
)

func getNytArticles() ([]models.Article, error) {
	articles := []models.Article{}
	rsp, err := NewYorkTimes()
	if err != nil {
		return nil, err
	}
	for _, a := range rsp.Results {
		articles = append(articles, models.Article{
			Title:       a.Title,
			URL:         a.URL,
			Thumbnail:   a.Thumbnail,
			PublishedAt: a.PublishedAt,
		})
	}
	return articles, nil
}

func getNewsArticles() ([]models.Article, error) {
	articles := []models.Article{}
	rsp, err := News()
	if err != nil {
		return nil, err
	}
	for _, a := range rsp.Results {
		articles = append(articles, models.Article{
			Title:       a.Title,
			URL:         a.URL,
			Thumbnail:   a.Thumbnail,
			PublishedAt: a.PublishedAt,
		})
	}
	return articles, nil
}

// Get returns an array of articles
func Get() *[]models.Article {
	nyt, err := getNytArticles()
	if err != nil {
		panic(err)
	}

	news, err := getNewsArticles()
	if err != nil {
		panic(err)
	}

	articles := append(nyt, news...)

	return &articles
}
