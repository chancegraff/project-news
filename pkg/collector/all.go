package collector

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
)

var apiURL = "/api/v1/ranks/articles"

// Vote ...
type Vote struct {
	ArticleID string
	Count     string
}

func all(w http.ResponseWriter, r *http.Request) {
	logger := utils.NewHTTPLogger("All", &w)

	offset := r.FormValue("offset")
	if offset == "" {
		offset = "0"
	}

	// Get articles
	var articles []models.Article
	store.Offset(offset).Limit(20).Order("published_at desc").Find(&articles)

	// Pick article IDs
	var articleIDs []string
	for _, art := range articles {
		articleIDs = append(articleIDs, fmt.Sprint(art.ID))
	}

	// Get ranks for picked IDs
	var ranks []Vote
	err := store.Select("article_id,count(*) as count").Where("article_id IN (?)", articleIDs).Where("created_at > ?", time.Now().AddDate(0, 0, -1)).Group("article_id").Find(&ranks).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
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
	logger.Okay(articles)
}
