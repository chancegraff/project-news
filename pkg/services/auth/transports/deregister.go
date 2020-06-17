package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/auth"
)

// DeregisterRequest ...
type DeregisterRequest struct {
	UserID string `json:"user"`
}

// DeregisterResponse ...
type DeregisterResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeDeregisterHTTPRequest ...
func DecodeDeregisterHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeregisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeDeregisterRPCRequest ...
func DecodeDeregisterRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeregisterRequest)
	return DeregisterRequest{UserID: req.UserID}, nil
}

// EncodeDeregisterRPCResponse ...
func EncodeDeregisterRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(DeregisterResponse)
	user := &pb.User{
		Email:      res.User.Email,
		Password:   res.User.Password,
		VerifiedAt: res.User.VerifiedAt.String(),
		Id:         int32(res.User.ID),
		CreatedAt:  res.User.CreatedAt.String(),
		UpdatedAt:  res.User.UpdatedAt.String(),
	}
	return &pb.DeregisterResponse{User: user, Err: res.Err}, nil
}
