package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/interfaces"
)

// DecodeVoteRequest ...
func DecodeVoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request interfaces.VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
