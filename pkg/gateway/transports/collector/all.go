package collector

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	pbc "github.com/chancegraff/project-news/api/collector"
)

// DecodeAllRequest ...
func DecodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	s, _ := ioutil.ReadAll(r.Body)
	var request *pbc.AllRequest
	log.Println("Decode all request", s)
	if err := json.NewDecoder(r.Body).Decode(request); err != io.EOF && err != nil {
		log.Println("Decode all request error", r)
		return nil, err
	}
	return request, nil
}

// EncodeAllResponse ...
func EncodeAllResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("Encode all respomse")
	return json.NewEncoder(w).Encode(response)
}
