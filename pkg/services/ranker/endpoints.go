package ranker

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeArticlesEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(articlesRequest)
		articles, err := svc.Articles(req.ArticleIDs)
		if err != nil {
			return articlesResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return articlesResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

func makeUserEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(userRequest)
		user, err := svc.User(req.UserID)
		if err != nil {
			return userResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return userResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

func makeVoteEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(voteRequest)
		article, err := svc.Vote(req.ArticleID, req.UserID)
		if err != nil {
			return voteResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return voteResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}
