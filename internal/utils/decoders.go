package utils

import (
	"encoding/json"
	"net/http"
)

// RequestBodyDecoder returns the body from a request
func RequestBodyDecoder(i interface{}, rq *http.Request) error {
	err := json.NewDecoder(rq.Body).Decode(&i)
	if err != nil {
		return err
	}
	return nil
}
