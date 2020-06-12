package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/models"
)

type userPayload struct {
	UserID int32 `json:"user"`
}

func user(wt http.ResponseWriter, rq *http.Request) {
	logger := utils.NewHTTPLogger("Verify", &wt)

	// Decode payload
	var pld userPayload
	err := json.NewDecoder(rq.Body).Decode(&pld)
	if err != nil {
		logger.Error(err, http.StatusBadRequest)
		return
	}

	// Get user votes
	var vts []models.Vote
	err = store.Select("article_id").Where("user_id = ?", fmt.Sprint(pld.UserID)).Find(&vts).Error
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	logger.Okay(&vts)
}
