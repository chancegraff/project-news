package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chancegraff/project-news/internal/utils"
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

	// Get array of article IDs by user ID
	vts, err := VotesByUser(fmt.Sprint(pld.UserID))
	if err != nil {
		logger.Error(err, http.StatusInternalServerError)
		return
	}

	logger.Okay(&vts)
}
