package ranker

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// DecodeVoteRequest ...
func DecodeVoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pbr.VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeVoteResponse ...
func EncodeVoteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
