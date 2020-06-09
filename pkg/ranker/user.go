package rest

import (
	"encoding/json"
	"net/http"

	"github.com/chancegraff/goutils/loggers"
	"github.com/chancegraff/project-news/pkg/models"
)

type payload struct {
	UserID string `json:"user"`
}

// user takes a user ID and returns an array of article IDs associated with it
func user(wt http.ResponseWriter, rq *http.Request) {
	logger := loggers.NewHTTPLogger("User", &wt)

	var pd payload
	err := json.NewDecoder(rq.Body).Decode(&pd)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	var vts []models.Vote
	err = store.Select("article_id").Where("user_id = ?", pd.UserID).Find(&vts).Error
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	var ids []string
	for _, vt := range vts {
		ids = append(ids, vt.ArticleID)
	}

	logger.Okay(ids)
}
