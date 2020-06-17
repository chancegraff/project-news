package ranker

import (
	"context"
	"errors"

	pbr "github.com/chancegraff/project-news/api/ranker"
)

// Vote ...
func (p *proxy) Vote(articleID, userID string) (*pbr.ArticleVotes, error) {
	voteResp, err := p.Client.Vote(context.Background(), &pbr.VoteRequest{ArticleID: articleID, UserID: userID})
	if err != nil {
		return nil, err
	}
	if err := voteResp.Err; err != "" {
		return nil, errors.New(voteResp.Err)
	}
	return voteResp.Article, nil
}
