package service

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/ranker"
)

// All will return articles from the database with their rank
func (s *service) All(offset int) ([]models.Article, error) {
	log.Println("All service called")

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
	response, err := s.RankerService.Articles(context.Background(), &pb.ArticlesRequest{ArticleIDs: articleIDs})
	if err != nil {
		return nil, err
	}
	articleVotes := response.Articles

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
