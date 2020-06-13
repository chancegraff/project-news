package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/ranker/interfaces"
)

// DecodeArticlesRequest ...
func DecodeArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request interfaces.ArticlesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
