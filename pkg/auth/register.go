package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chancegraff/goutils/loggers"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func register(wt http.ResponseWriter, rq *http.Request) {
	logger := loggers.NewHTTPLogger("Register", &wt)

	// Decode body credentials
	var cred models.Credentials
	err := json.NewDecoder(rq.Body).Decode(&cred)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Search for existing user
	var tmp models.User
	err = store.Where("email = ?", cred.Email).First(&tmp).Error
	if gorm.IsRecordNotFoundError(err) == false {
		logger.Error(err, http.StatusConflict)
		return
	}

	// Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(cred.Password), 8)
	user := models.User{
		Credentials: models.Credentials{
			Email:    cred.Email,
			Password: string(hash),
		},
	}

	// Create user
	store.Create(&user)

	// Return result
	logger.Okay(&models.UserWithArticles{ID: fmt.Sprint(user.ID), Email: user.Email})
}
