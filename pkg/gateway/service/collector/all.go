package collector

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// All ...
func (s *service) All(offset int) ([]*pbc.Article, error) {
	// Call collector service
	articles, err := s.Proxy.Collector.All(offset)
	if err != nil {
		return nil, err
	}
	log.Println("Service all 1")
	// Pick article IDs
	var articleIDs []string
	for _, art := range articles {
		articleIDs = append(articleIDs, fmt.Sprint(art.Id))
	}
	log.Println("Service all 2")
	// Call ranker service
	articleVotes, err := s.Proxy.Ranker.Articles(articleIDs)
	if err != nil {
		return nil, err
	}
	log.Println("Service all 3")
	// Put articles into order
	sort.Slice(articles, func(i, j int) bool {
		iRank, jRank := "0", "0"
		iArticle, jArticle := articles[i], articles[j]

		for _, articleVote := range articleVotes {
			if articleVote.ArticleID == fmt.Sprint(iArticle.Id) {
				iRank = fmt.Sprint(articleVote.Votes)
			} else if articleVote.ArticleID == fmt.Sprint(jArticle.Id) {
				jRank = fmt.Sprint(articleVote.Votes)
			}
		}

		iRankInt, _ := strconv.ParseInt(iRank, 10, 32)
		jRankInt, _ := strconv.ParseInt(jRank, 10, 32)

		return iRankInt > jRankInt
	})
	log.Println("Service all 4")
	// Return response
	return articles, nil
}
