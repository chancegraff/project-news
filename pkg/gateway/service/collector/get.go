package collector

import (
	pbc "github.com/chancegraff/project-news/api/collector"
)

// Get ...
func (s *service) Get(id int) (*pbc.Article, error) {
	return s.Proxy.Collector.Get(id)
}
