package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/ranker"
)

// UserRequest ...
type UserRequest struct {
	UserID string `json:"user"`
}

// UserResponse ...
type UserResponse struct {
	User models.UserVotes `json:"user"`
	Err  string           `json:"err,omitempty"`
}

// DecodeUserHTTPRequest ...
func DecodeUserHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeUserRPCRequest ...
func DecodeUserRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return UserRequest{UserID: req.UserID}, nil
}

// EncodeUserRPCResponse ...
func EncodeUserRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(UserResponse)
	user := &pb.UserVotes{
		UserID: res.User.UserID,
		Votes:  make([]*pb.Vote, len(res.User.Votes)),
	}
	for i := range res.User.Votes {
		user.Votes[i] = &pb.Vote{
			UserID:    res.User.Votes[i].UserID,
			ArticleID: res.User.Votes[i].ArticleID,
			Id:        int32(res.User.Votes[i].ID),
			CreatedAt: res.User.Votes[i].CreatedAt.String(),
			UpdatedAt: res.User.Votes[i].UpdatedAt.String(),
		}
	}
	return &pb.UserResponse{User: user, Err: res.Err}, nil
}
