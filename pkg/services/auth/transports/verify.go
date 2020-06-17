package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/auth"
)

// VerifyRequest ...
type VerifyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// VerifyResponse ...
type VerifyResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeVerifyHTTPRequest ...
func DecodeVerifyHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeVerifyRPCRequest ...
func DecodeVerifyRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VerifyRequest)
	return VerifyRequest{Email: req.Email, Password: req.Password}, nil
}

// EncodeVerifyRPCResponse ...
func EncodeVerifyRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(VerifyResponse)
	user := &pb.User{
		Email:      res.User.Email,
		Password:   res.User.Password,
		VerifiedAt: res.User.VerifiedAt.String(),
		Id:         int32(res.User.ID),
		CreatedAt:  res.User.CreatedAt.String(),
		UpdatedAt:  res.User.UpdatedAt.String(),
	}
	return &pb.VerifyResponse{User: user, Err: res.Err}, nil
}
