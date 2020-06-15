package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/collector"
)

// AllRequest ...
type AllRequest struct {
	Offset int `json:"offset,omitempty"`
}

// AllResponse ...
type AllResponse struct {
	Articles []models.Article `json:"articles"`
	Err      string           `json:"err,omitempty"`
}

// DecodeAllHTTPRequest ...
func DecodeAllHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeAllRPCRequest ...
func DecodeAllRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AllRequest)
	return AllRequest{Offset: int(req.Offset)}, nil
}

// EncodeAllRPCResponse ...
func EncodeAllRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AllResponse)
	articles := make([]*pb.Article, len(res.Articles))
	for i := range res.Articles {
		articles[i] = &pb.Article{
			Title:       res.Articles[i].Title,
			Url:         res.Articles[i].URL,
			Thumbnail:   res.Articles[i].Thumbnail,
			PublishedAt: res.Articles[i].PublishedAt.String(),
			Id:          int32(res.Articles[i].ID),
			CreatedAt:   res.Articles[i].CreatedAt.String(),
			UpdatedAt:   res.Articles[i].UpdatedAt.String(),
		}
	}
	return &pb.AllResponse{Articles: articles, Err: res.Err}, nil
}
