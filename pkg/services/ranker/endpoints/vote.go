package endpoints

import (
	"context"

	"github.com/chancegraff/project-news/pkg/services/ranker/interfaces"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/go-kit/kit/endpoint"
)

func makeVoteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(interfaces.VoteRequest)
		article, err := svc.Vote(req.ArticleID, req.UserID)
		if err != nil {
			return interfaces.VoteResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return interfaces.VoteResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}
