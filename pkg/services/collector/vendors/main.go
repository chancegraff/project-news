package vendors

import (
	"context"
	"errors"
	"time"

	"github.com/chancegraff/project-news/pkg/services/collector/manager"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Server will routinely push new articles into the database
type Server interface {
	Start(parent context.Context, lgr log.Logger) error
	Run(parent context.Context, lgr log.Logger) error
	Stop(parent context.Context, lgr log.Logger) error
}

type server struct {
	manager *manager.Manager
}

// Start will run the server and block until it returns
func (s *server) Start(parent context.Context, lgr log.Logger) error {
	_, cancel := context.WithCancel(parent)
	level.Info(lgr).Log("msg", "vendor collector started")
	err := s.Run(parent, lgr)
	cancel()
	return err
}

// Run will periodically collect articles into the database until the context closes
func (s *server) Run(parent context.Context, lgr log.Logger) error {
	for {
		select {
		case <-parent.Done():
			return errors.New("Vendors collector shutting down")
		default:
			level.Info(lgr).Log("msg", "vendor collector running")
			s.Collect(parent)
			time.Sleep(5 * time.Minute)
		}
	}
}

func (s *server) Stop(parent context.Context, lgr log.Logger) error {
	select {
	case <-parent.Done():
		level.Info(lgr).Log("msg", "vendor collector stopped")
		return parent.Err()
	}
}

// NewServer ...
func NewServer(manager *manager.Manager) Server {
	return &server{
		manager: manager,
	}
}
