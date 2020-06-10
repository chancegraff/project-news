package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
)

var apiURL = "/api/v1/ranks/articles"

type rank struct {
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

	// Marshal IDs into JSON
	js, err := json.Marshal(articleIDs)
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	// Post to endpoint
	url := fmt.Sprint(r.URL.Scheme, r.URL.Host, apiURL)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(js))
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	// Decode into array
	var ranks []rank
	err = json.NewDecoder(res.Body).Decode(&ranks)
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
