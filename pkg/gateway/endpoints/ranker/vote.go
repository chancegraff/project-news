package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"

	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/endpoint"
)

// MakeVoteEndpoint ...
func MakeVoteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbr.VoteRequest)
		article, err := svc.Ranker.Vote(req.ArticleID, req.UserID)
		if err != nil {
			return pbr.VoteResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return pbr.VoteResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

// Vote ...
func (e Endpoints) Vote(ctx context.Context, articleID, userID string) (*pbr.ArticleVotes, error) {
	req := &pbr.VoteRequest{ArticleID: articleID, UserID: userID}
	resp, err := e.VoteEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	voteResp := resp.(*pbr.VoteResponse)
	if voteResp.Err != "" {
		return nil, errors.New(voteResp.Err)
	}
	return voteResp.Article, nil
}
