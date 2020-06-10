package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var apiURL = "/api/v1/ranks/user"

// VotesByUser takes a user ID and returns an array of article IDs associated with it
func VotesByUser(userID string, r *http.Request) (*[]string, error) {
	// Marshal user ID into JSON
	js, err := json.Marshal(map[string]string{"user": userID})
	if err != nil {
		return nil, err
	}

	// Post to endpoint
	url := fmt.Sprint(r.URL.Scheme, r.URL.Host, apiURL)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(js))
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
