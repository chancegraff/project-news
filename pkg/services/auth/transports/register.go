package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/auth"
)

// RegisterRequest ...
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse ...
type RegisterResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeRegisterHTTPRequest ...
func DecodeRegisterHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeRegisterRPCRequest ...
func DecodeRegisterRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterRequest)
	return RegisterRequest{Email: req.Email, Password: req.Password}, nil
}

// EncodeRegisterRPCResponse ...
func EncodeRegisterRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RegisterResponse)
	user := &pb.User{
		Email:      res.User.Email,
		Password:   res.User.Password,
		VerifiedAt: res.User.VerifiedAt.String(),
		Id:         int32(res.User.ID),
		CreatedAt:  res.User.CreatedAt.String(),
		UpdatedAt:  res.User.UpdatedAt.String(),
	}
	return &pb.RegisterResponse{User: user, Err: res.Err}, nil
}
