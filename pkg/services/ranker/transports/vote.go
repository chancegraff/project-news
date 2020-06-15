package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/ranker"
)

// VoteRequest ...
type VoteRequest struct {
	ArticleID string `json:"article"`
	UserID    string `json:"user"`
}

// VoteResponse ...
type VoteResponse struct {
	Article models.ArticleVotes `json:"article"`
	Err     string              `json:"err,omitempty"`
}

// DecodeVoteHTTPRequest ...
func DecodeVoteHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeVoteRPCRequest ...
func DecodeVoteRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VoteRequest)
	return VoteRequest{UserID: req.UserID, ArticleID: req.ArticleID}, nil
}

// EncodeVoteRPCResponse ...
func EncodeVoteRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(VoteResponse)
	article := &pb.ArticleVotes{
		ArticleID: res.Article.ArticleID,
		Votes:     int32(res.Article.Votes),
	}
	return &pb.VoteResponse{Article: article, Err: res.Err}, nil
}
