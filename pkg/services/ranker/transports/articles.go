package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/ranker"
)

// ArticlesRequest ...
type ArticlesRequest struct {
	ArticleIDs []string `json:"articles"`
}

// ArticlesResponse ...
type ArticlesResponse struct {
	Articles []models.ArticleVotes `json:"articles"`
	Err      string                `json:"err,omitempty"`
}

// DecodeArticlesHTTPRequest ...
func DecodeArticlesHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ArticlesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeArticlesRPCRequest ...
func DecodeArticlesRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ArticlesRequest)
	return ArticlesRequest{ArticleIDs: req.ArticleIDs}, nil
}

// EncodeArticlesRPCResponse ...
func EncodeArticlesRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ArticlesResponse)
	articles := make([]*pb.ArticleVotes, len(res.Articles))
	for i := range res.Articles {
		articles[i] = &pb.ArticleVotes{
			ArticleID: res.Articles[i].ArticleID,
			Votes:     int32(res.Articles[i].Votes),
		}
	}
	return &pb.ArticlesResponse{Articles: articles, Err: res.Err}, nil
}
