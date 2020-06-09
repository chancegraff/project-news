package rest

import (
	"encoding/json"
	"net/http"

	"github.com/chancegraff/goutils/loggers"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/jinzhu/gorm"
)

// vote takes an article ID and a user ID and returns a success code
func vote(wt http.ResponseWriter, rq *http.Request) {
	logger := loggers.NewHTTPLogger("Vote", &wt)

	// Decode body into model
	var vt, tmp models.Vote
	err := json.NewDecoder(rq.Body).Decode(&vt)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Create for new and delete for old
	err = store.Where(models.Vote{UserID: vt.UserID, ArticleID: vt.ArticleID}).First(&tmp).Error
	if gorm.IsRecordNotFoundError(err) {
		store.Create(&vt)
	} else {
		store.Delete(&tmp)
	}

	// Return array of user IDs associated with article
	var vts []models.Vote
	var vtIDs []string
	err = store.Select("user_id").Where("article_id = ?", vt.ArticleID).Find(&vts).Pluck("user_id", &vtIDs).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	logger.Okay(vtIDs)
}
