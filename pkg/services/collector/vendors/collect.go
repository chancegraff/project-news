package vendors

import (
	"context"

	"github.com/chancegraff/project-news/internal/vendors"
)

// Collect ...
func (c *server) Collect(parent context.Context) {
	_, cancel := context.WithCancel(parent)
	select {
	case <-parent.Done():
		cancel()
		return
	default:
		articles := vendors.Get()
		c.manager.Batch(articles)
		cancel()
		return
	}
}
