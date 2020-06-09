package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chancegraff/goutils/loggers"
	"github.com/chancegraff/project-news/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func verify(wt http.ResponseWriter, rq *http.Request) {
	logger := loggers.NewHTTPLogger("Verify", &wt)

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

	// Get array of article IDs by user ID
	vts, err := VotesByUser(fmt.Sprint(usr.ID))
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	logger.Okay(&models.UserWithArticles{ID: fmt.Sprint(usr.ID), Email: usr.Email, Articles: *vts})
}
