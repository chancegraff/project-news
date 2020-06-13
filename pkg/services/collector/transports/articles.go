package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chancegraff/project-news/pkg/services/collector/interfaces"
)

// DecodeArticlesResponse ...
func DecodeArticlesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response interfaces.ArticlesResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
