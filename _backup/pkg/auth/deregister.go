package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/jinzhu/gorm"
)

type deregisterPayload struct {
	UserID string `json:"user"`
}

func deregister(wt http.ResponseWriter, rq *http.Request) {
	logger := utils.NewHTTPLogger("Deregister", &wt)

	// Decode payload
	var pld deregisterPayload
	err := json.NewDecoder(rq.Body).Decode(&pld)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Search for existing user
	var tmp models.User
	err = store.First(&tmp, pld.UserID).Error
	if gorm.IsRecordNotFoundError(err) {
		logger.Error(err, http.StatusNotFound)
		return
	}

	// Deregister current user token
	store.Model(&tmp).Update("verified_at", time.Now())

	// Return result
	logger.Okay(map[string]string{"status": "ok"})
}
