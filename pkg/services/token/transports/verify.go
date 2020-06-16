package transports

import (
	"context"

	pb "github.com/chancegraff/project-news/api/token"
	"github.com/chancegraff/project-news/internal/models"
)

// VerifyRequest ...
type VerifyRequest struct {
	Identifiers models.Identifiers `json:"identifiers"`
	Client      models.Client      `json:"client"`
}

// VerifyResponse ...
type VerifyResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

// DecodeVerifyRPCRequest ...
func DecodeVerifyRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VerifyRequest)
	return VerifyRequest{
		Client:      DecodeProtoClient(req.Client),
		Identifiers: DecodeProtoIdentifiers(req.Identifiers),
	}, nil
}

// EncodeVerifyRPCResponse ...
func EncodeVerifyRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(VerifyResponse)
	return &pb.VerifyResponse{Hash: res.Hash, Err: res.Err}, nil
}
