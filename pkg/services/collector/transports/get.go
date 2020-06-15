package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/collector"
)

// GetRequest ...
type GetRequest struct {
	ID int `json:"id"`
}

// GetResponse ...
type GetResponse struct {
	Article models.Article `json:"article"`
	Err     string         `json:"err,omitempty"`
}

// DecodeGetHTTPRequest ...
func DecodeGetHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeGetRPCRequest ...
func DecodeGetRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRequest)
	return GetRequest{ID: int(req.Id)}, nil
}

// EncodeGetRPCResponse ...
func EncodeGetRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetResponse)
	article := &pb.Article{
		Title:       res.Article.Title,
		Url:         res.Article.URL,
		Thumbnail:   res.Article.Thumbnail,
		PublishedAt: res.Article.PublishedAt.String(),
		Id:          int32(res.Article.ID),
		CreatedAt:   res.Article.CreatedAt.String(),
		UpdatedAt:   res.Article.UpdatedAt.String(),
	}
	return &pb.GetResponse{Article: article, Err: res.Err}, nil
}
