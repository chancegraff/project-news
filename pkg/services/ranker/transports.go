package ranker

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request articlesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request userRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeVoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request voteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
