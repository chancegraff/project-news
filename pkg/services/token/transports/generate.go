package transports

import (
	"context"

	"github.com/chancegraff/project-news/internal/models"

	pb "github.com/chancegraff/project-news/api/token"
)

// GenerateRequest ...
type GenerateRequest struct {
	Identifiers models.Identifiers `json:"identifiers"`
	Client      models.Client      `json:"client"`
}

// GenerateResponse ...
type GenerateResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

// DecodeGenerateRPCRequest ...
func DecodeGenerateRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GenerateRequest)
	return GenerateRequest{
		Client:      DecodeProtoClient(req.Client),
		Identifiers: DecodeProtoIdentifiers(req.Identifiers),
	}, nil
}

// EncodeGenerateRPCResponse ...
func EncodeGenerateRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GenerateResponse)
	return &pb.GenerateResponse{Hash: res.Hash, Err: res.Err}, nil
}
