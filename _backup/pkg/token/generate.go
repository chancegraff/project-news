package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// TODO Fix creating new tokens for existing identifiers

func generate(wt http.ResponseWriter, rq *http.Request) {
	logger := utils.NewHTTPLogger("Generate", &wt)

	// Decode body
	identifiers := models.Identifiers{}
	err := utils.RequestBodyDecoder(&identifiers, rq)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Build payload
	payload, err := utils.GetJSONHashFromStruct(identifiers)
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	// Build secret
	secret := uuid.New()

	// Encode with hmac sha256
	hash := utils.HMACSHA256([]byte(*payload), secret.String())

	// Check for existing hash
	client := &models.Client{Hash: hash, Secret: secret, ExpiredAt: time.Now().AddDate(0, 0, 1)}
	err = store.Where("hash = ? AND expired_at > ?", client.Hash, time.Now()).First(&models.Client{}).Error
	if gorm.IsRecordNotFoundError(err) == false {
		logger.Error(errors.New("hash already exists"), http.StatusConflict)
		return
	}

	// Store to DB
	store.Create(client)

	// Return result
	logger.Okay(client.Hash)
}
