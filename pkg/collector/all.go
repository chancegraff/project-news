package collector

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/chancegraff/project-news/pkg/models"
)

type vote struct {
	ArticleID string
	Count     string
}

type bodyAll struct {
	Offset string
}

func all(data interface{}) ([]byte, error) {
	// Typecast into body
	bd := data.(bodyAll)

	// Get articles
	var articles []models.Article
	store.Offset(bd.Offset).Limit(20).Order("published_at desc").Find(&articles)

	// Pick article IDs
	var articleIDs []string
	for _, art := range articles {
		articleIDs = append(articleIDs, fmt.Sprint(art.ID))
	}

	// Get ranks for picked IDs
	var ranks []vote
	err := store.Select("article_id,count(*) as count").Where("article_id IN (?)", articleIDs).Where("created_at > ?", time.Now().AddDate(0, 0, -1)).Group("article_id").Find(&ranks).Error
	if err != nil {
		return nil, err
	}

	// Sort array
	sort.Slice(articles, func(i, j int) bool {
		iRank, jRank := "0", "0"
		iArticle, jArticle := articles[i], articles[j]

		for _, rank := range ranks {
			if rank.ArticleID == fmt.Sprint(iArticle.ID) {
				iRank = rank.Count
			} else if rank.ArticleID == fmt.Sprint(jArticle.ID) {
				jRank = rank.Count
			}
		}

		iRankInt, _ := strconv.ParseInt(iRank, 10, 32)
		jRankInt, _ := strconv.ParseInt(jRank, 10, 32)

		return iRankInt > jRankInt
	})

	return json.Marshal(articles)
}
