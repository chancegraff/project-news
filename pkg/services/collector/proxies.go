package collector

import (
	"errors"
	"net/url"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func articlesEndpointMiddleware(proxyURL string) ServiceMiddleware {
	return func(next Service) Service {
		return articlesmw{
			articles: makeArticlesProxy(proxyURL),
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

func makeArticlesProxy(proxyURL string) endpoint.Endpoint {
	u, err := url.Parse(proxyURL)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/articles"
	}
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeArticlesResponse,
	).Endpoint()
}
