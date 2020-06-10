package ranker

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chancegraff/goutils/loggers"
)

// Vote ...
type Vote struct {
	ArticleID string
	Count     string
}

// articles takes an array of article IDs and returns the count of Votes associated with each
func articles(wt http.ResponseWriter, rq *http.Request) {
	logger := loggers.NewHTTPLogger("Articles", &wt)

	var artIDs []string
	err := json.NewDecoder(rq.Body).Decode(&artIDs)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	var rsp []Vote
	err = store.Select("article_id,count(*) as count").Where("article_id IN (?)", artIDs).Where("created_at > ?", time.Now().AddDate(0, 0, -1)).Group("article_id").Find(&rsp).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	logger.Okay(rsp)
}
