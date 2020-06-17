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
func NewProxy(ctx context.Context) (*Proxy, error) {
	prx := Proxy{
		Auth:      auth.NewProxy(),
		Collector: collector.NewProxy(),
		Ranker:    ranker.NewProxy(),
		Token:     token.NewProxy(),
	}
	err := prx.Start(ctx)
	if err != nil {
		return nil, err
	}
	return &prx, nil
}

// Start will connect all of the proxy services or panic
func (p Proxy) Start(ctx context.Context) error {
	err := p.Auth.Start(ctx)
	if err != nil {
		return err
	}
	err = p.Collector.Start(ctx)
	if err != nil {
		return err
	}
	err = p.Ranker.Start(ctx)
	if err != nil {
		return err
	}
	err = p.Token.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Stop will disconnect all of the proxy services
func (p Proxy) Stop() {
	p.Auth.Stop()
	p.Collector.Stop()
	p.Ranker.Stop()
	p.Token.Stop()
}
