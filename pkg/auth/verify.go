package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func verify(wt http.ResponseWriter, rq *http.Request) {
	logger := utils.NewHTTPLogger("Verify", &wt)

	var cred models.Credentials
	err := json.NewDecoder(rq.Body).Decode(&cred)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	var usr models.User
	err = store.Where("email = ?", cred.Email).First(&usr).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(cred.Password))
	if err != nil {
		logger.Error(err, http.StatusUnauthorized)
		return
	}

	// Get array of votes
	var vts []models.Vote
	err = store.Select("article_id").Where("user_id = ?", fmt.Sprint(usr.ID)).Find(&vts).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	// Pull article IDs from votes
	var ids []string
	for _, vt := range vts {
		ids = append(ids, vt.ArticleID)
	}

	logger.Okay(&models.UserWithArticles{ID: fmt.Sprint(usr.ID), Email: usr.Email, Articles: ids})
}
