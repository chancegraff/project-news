package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/auth"
)

// UserRequest ...
type UserRequest struct {
	UserID string `json:"user"`
}

// UserResponse ...
type UserResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
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
	user := &pb.User{
		Email:      res.User.Email,
		Password:   res.User.Password,
		VerifiedAt: res.User.VerifiedAt.String(),
		Id:         int32(res.User.ID),
		CreatedAt:  res.User.CreatedAt.String(),
		UpdatedAt:  res.User.UpdatedAt.String(),
	}
	return &pb.UserResponse{User: user, Err: res.Err}, nil
}
