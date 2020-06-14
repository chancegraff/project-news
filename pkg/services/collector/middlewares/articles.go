package middlewares

import (
	"errors"
	"net/url"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/transports"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	httptransport "github.com/go-kit/kit/transport/http"
)

// ArticlesProxyMiddleware ...
func ArticlesProxyMiddleware() service.Middleware {
	queryRate := utils.GetQueryRateLimit()
	endpoint := MakeArticlesProxy()
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(endpoint)
	endpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), queryRate))(endpoint)
	return func(next service.Service) service.Service {
		return ArticlesMiddleware{
			Service:          next,
			ArticlesEndpoint: endpoint,
		}
	}
}

// MakeArticlesProxy ...
func MakeArticlesProxy() endpoint.Endpoint {
	rankerAddress := utils.GetRankerAddress()
	u, err := url.Parse(rankerAddress)
	if err != nil {
		panic(err)
	}
	return httptransport.NewClient(
		"GET",
		u,
		transports.EncodeRequest,
		transports.DecodeArticlesResponse,
	).Endpoint()
}

// ArticlesMiddleware ...
type ArticlesMiddleware struct {
	service.Service
	ArticlesEndpoint endpoint.Endpoint
}

// Articles ...
func (mw ArticlesMiddleware) Articles(articleIDs []string) ([]models.ArticleVotes, error) {
	response, err := mw.ArticlesEndpoint(nil, transports.ArticlesRequest{ArticleIDs: articleIDs})
	if err != nil {
		return []models.ArticleVotes{}, err
	}
	resp := response.(transports.ArticlesResponse)
	if resp.Err != "" {
		return resp.Articles, errors.New(resp.Err)
	}
	return resp.Articles, nil
}
