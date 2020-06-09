package rest

import (
	"net/http"

	"github.com/chancegraff/goutils/loggers"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	logger := loggers.NewHTTPLogger("Get", &w)

	articleID := mux.Vars(r)["id"]

	var article models.Article
	err := store.First(&article, articleID).Error
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	logger.Okay(article)
}
