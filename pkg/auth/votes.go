package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var apiURL = "http://localhost:8001/api/v1/ranks/user"

// VotesByUser takes a user ID and returns an array of article IDs associated with it
func VotesByUser(userID string) (*[]string, error) {
	// Marshal user ID into JSON
	js, err := json.Marshal(map[string]string{"user": userID})
	if err != nil {
		return nil, err
	}

	// Post to endpoint
	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	// Decode into array of IDs
	var arts []string
	err = json.NewDecoder(res.Body).Decode(&arts)
	if err != nil {
		return nil, err
	}

	return &arts, nil
}
