package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	requestURL, _ := url.Parse(r.RequestURI)
	address := fmt.Sprint(requestURL.Scheme, requestURL.Host, apiURL)
	return nil, errors.New(address)
	// res, err := http.Post(address, "application/json", bytes.NewBuffer(js))
	// if err != nil {
	// 	return nil, err
	// }

	// // Decode into array of IDs
	// var arts []string
	// err = json.NewDecoder(res.Body).Decode(&arts)
	// if err != nil {
	// 	return nil, err
	// }

	// return &arts, nil
}
