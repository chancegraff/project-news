package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/chancegraff/project-news/pkg/services/ranker/transports"
	"github.com/go-kit/kit/endpoint"
)

// MakeVoteEndpoint ...
func MakeVoteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.VoteRequest)
		article, err := svc.Vote(req.ArticleID, req.UserID)
		if err != nil {
			return transports.VoteResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return transports.VoteResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}
