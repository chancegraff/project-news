package collector

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/models"
	"github.com/go-kit/kit/endpoint"
)

// Service implements the collector interface
type Service interface {
	All(offset int) ([]models.Article, error)
	Get(id int) (models.Article, error)
}

type service struct {
	Manager  *Manager
	articles endpoint.Endpoint
}

// Get will find and return a single article from the database that matches the ID
func (s *service) Get(id int) (models.Article, error) {
	return s.Manager.First(id)
}

// All will return articles from the database with their rank
func (s *service) All(offset int) ([]models.Article, error) {
	// Get articles
	articles, err := s.Manager.List(offset, 20)
	if err != nil {
		return nil, err
	}

	// Pick article IDs
	var articleIDs []string
	for _, art := range articles {
		articleIDs = append(articleIDs, fmt.Sprint(art.ID))
	}

	// Call ranker service
	response, err := s.articles(nil, articlesRequest{ArticleIDs: articleIDs})
	if err != nil {
		return nil, err
	}
	articleVotes := response.(articlesResponse).Articles

	// Put articles into order
	sort.Slice(articles, func(i, j int) bool {
		iRank, jRank := "0", "0"
		iArticle, jArticle := articles[i], articles[j]

		for _, articleVote := range articleVotes {
			if articleVote.ArticleID == fmt.Sprint(iArticle.ID) {
				iRank = fmt.Sprint(articleVote.Votes)
			} else if articleVote.ArticleID == fmt.Sprint(jArticle.ID) {
				jRank = fmt.Sprint(articleVote.Votes)
			}
		}

		iRankInt, _ := strconv.ParseInt(iRank, 10, 32)
		jRankInt, _ := strconv.ParseInt(jRank, 10, 32)

		return iRankInt > jRankInt
	})

	// Return response
	return articles, nil
}

// NewService instantiates the service with a connection to the database
func newService() (*service, error) {
	store, err := db.NewStore()
	if err != nil {
		return nil, err
	}
	return &service{Manager: &Manager{store}}, nil
}

// ServiceMiddleware is a chainable middleware for Service
type ServiceMiddleware func(Service) Service
