package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
)

type verifyPayload struct {
	Identifiers models.Identifiers `json:"identifiers"`
	UserID      string             `json:"user"`
	Hash        string             `json:"hash"`
}

func verify(wt http.ResponseWriter, rq *http.Request) {
	logger := utils.NewHTTPLogger("Generate", &wt)

	// Decode body
	payload := verifyPayload{}
	err := utils.RequestBodyDecoder(&payload, rq)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Retrieve client
	client := &models.Client{Hash: payload.Hash, UserID: payload.UserID}
	err = store.Where(*client).Where("expired_at > ?", time.Now()).Find(client).Error
	if err != nil {
		logger.Error(err, http.StatusNotFound)
		return
	}

	// Generate hash from payload identifiers
	identifiers, err := utils.GetJSONHashFromStruct(payload.Identifiers)
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	// Encode identifiers hash
	hash := utils.HMACSHA256([]byte(*identifiers), client.Secret.String())

	// Verify hashes match
	if hash != client.Hash {
		logger.Error(errors.New("invalid"), http.StatusForbidden)
		return
	}

	// Return success
	logger.Okay(map[string]string{"status": "ok"})
}
