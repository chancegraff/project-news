package endpoints

import (
	"context"
	"errors"

	"github.com/chancegraff/project-news/internal/models"
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

// Vote ...
func (e Endpoints) Vote(ctx context.Context, articleID string) (models.ArticleVotes, error) {
	req := transports.VoteRequest{ArticleID: articleID}
	resp, err := e.VoteEndpoint(ctx, req)
	if err != nil {
		return models.ArticleVotes{}, err
	}
	voteResp := resp.(transports.VoteResponse)
	if voteResp.Err != "" {
		return voteResp.Article, errors.New(voteResp.Err)
	}
	return voteResp.Article, nil
}
