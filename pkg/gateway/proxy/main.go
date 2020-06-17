package proxy

import (
	"context"

	"github.com/chancegraff/project-news/pkg/gateway/proxy/auth"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/collector"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/ranker"
	"github.com/chancegraff/project-news/pkg/gateway/proxy/token"
	"github.com/go-kit/kit/log"
)

// Proxy is the root level container for the proxy
type Proxy struct {
	Auth      auth.Proxy
	Collector collector.Proxy
	Ranker    ranker.Proxy
	Token     token.Proxy
}

// NewProxy will create proxy connections to the various services
func NewProxy(ctx context.Context, lgr log.Logger) (*Proxy, error) {
	prx := Proxy{
		Auth:      auth.NewProxy(),
		Collector: collector.NewProxy(),
		Ranker:    ranker.NewProxy(),
		Token:     token.NewProxy(),
	}
	err := prx.Start(ctx, lgr)
	if err != nil {
		return nil, err
	}
	return &prx, nil
}

// Start will connect all of the proxy services or panic
func (p Proxy) Start(ctx context.Context, lgr log.Logger) error {
	err := p.Auth.Start(ctx)
	if err != nil {
		return err
	}
	lgr.Log("package", "proxy", "service", "auth", "event", "started")
	err = p.Collector.Start(ctx)
	if err != nil {
		return err
	}
	lgr.Log("package", "proxy", "service", "collector", "event", "started")
	err = p.Ranker.Start(ctx)
	if err != nil {
		return err
	}
	lgr.Log("package", "proxy", "service", "ranker", "event", "started")
	err = p.Token.Start(ctx)
	if err != nil {
		return err
	}
	lgr.Log("package", "proxy", "service", "token", "event", "started")
	return nil
}

// Stop will disconnect all of the proxy services
func (p Proxy) Stop() {
	p.Auth.Stop()
	p.Collector.Stop()
	p.Ranker.Stop()
	p.Token.Stop()
}
