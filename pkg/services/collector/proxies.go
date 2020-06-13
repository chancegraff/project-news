package collector

import (
	"errors"
	"net/url"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func articlesEndpointMiddleware() ServiceMiddleware {
	return func(next Service) Service {
		return articlesmw{
			Service:  next,
			articles: makeArticlesEndpoint(),
		}
	}
}

type articlesmw struct {
	Service
	articles endpoint.Endpoint
}

func (mw articlesmw) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	response, err := mw.articles(nil, articlesRequest{articleIDs})
	if err != nil {
		return []models.ArticleVotes{}, err
	}
	resp := response.(articlesResponse)
	if resp.Err != "" {
		return resp.Articles, errors.New(resp.Err)
	}
	return resp.Articles, nil
}

func makeArticlesEndpoint() endpoint.Endpoint {
	u, err := url.Parse("http://ranker.project-news-voter.app.localspace:7998/")
	if err != nil {
		panic(err)
	}
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeArticlesResponse,
	).Endpoint()
}
