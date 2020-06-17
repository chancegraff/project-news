package proxy

import (
	"context"

	"github.com/chancegraff/project-news/pkg/gateway/proxy/auth"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/collector"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/token"
)

// Proxy is the root level container for the proxy
type Proxy struct {
	Auth      auth.Proxy
	Collector collector.Proxy
	Ranker    ranker.Proxy
	Token     token.Proxy
}

// NewProxy will create proxy connections to the various services
func NewProxy() Proxy {
	return Proxy{
		Auth:      auth.NewProxy(),
		Collector: collector.NewProxy(),
		Ranker:    ranker.NewProxy(),
		Token:     token.NewProxy(),
	}
}

// Start will connect all of the proxy services or panic
func (p Proxy) Start(ctx context.Context) {
	MustStart(ctx, p.Auth.Start)
	MustStart(ctx, p.Collector.Start)
	MustStart(ctx, p.Ranker.Start)
	MustStart(ctx, p.Token.Start)
}

// MustStart will call the function or panic if it errors
func MustStart(ctx context.Context, start func(ctx context.Context) error) {
	err := start(ctx)
	if err != nil {
		panic(err)
	}
}

// Stop will disconnect all of the proxy services
func (p Proxy) Stop() {
	p.Auth.Stop()
	p.Collector.Stop()
	p.Ranker.Stop()
	p.Token.Stop()
}
