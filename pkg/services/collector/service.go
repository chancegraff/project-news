package collector

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/models"
)

// Service implements the collector interface
type Service interface {
	All(offset int) ([]models.Article, error)
	Get(id int) (models.Article, error)
}

type service struct{}

// Get will return an article from the database
func (service) Get(id int) (models.Article, error) {
	// Instantiate a database connection
	dbService, _ := db.NewService()
	defer dbService.Stop()

	// Get article
	var article models.Article
	err := dbService.Store.Database.First(&article, id).Error
	if err != nil {
		return article, err
	}

	// Return result
	return article, nil
}

// All will return articles from the database with their rank
func (service) All(offset int) ([]models.Article, error) {
	// Instantiate a database connection
	dbService, _ := db.NewService()
	defer dbService.Stop()

	// Get articles
	articles, err := dbService.Articles.List(offset, 20)
	if err != nil {
		return nil, err
	}

	// Pick article IDs
	var articleIDs []string
	for _, art := range articles {
		articleIDs = append(articleIDs, fmt.Sprint(art.ID))
	}

	// TODO Make call to ranker service here

	// Get ranks for picked IDs
	type vote struct {
		ArticleID string
		Count     string
	}

	var ranks []vote
	err = dbService.Store.Database.Select("article_id,count(*) as count").Where("article_id IN (?)", articleIDs).Where("created_at > ?", time.Now().AddDate(0, 0, -1)).Group("article_id").Find(&ranks).Error
	if err != nil {
		return nil, err
	}

	// Sort array
	sort.Slice(articles, func(i, j int) bool {
		iRank, jRank := "0", "0"
		iArticle, jArticle := articles[i], articles[j]

		for _, r := range ranks {
			if r.ArticleID == fmt.Sprint(iArticle.ID) {
				iRank = r.Count
			} else if r.ArticleID == fmt.Sprint(jArticle.ID) {
				jRank = r.Count
			}
		}

		iRankInt, _ := strconv.ParseInt(iRank, 10, 32)
		jRankInt, _ := strconv.ParseInt(jRank, 10, 32)

		return iRankInt > jRankInt
	})

	// Return response
	return articles, nil
}
