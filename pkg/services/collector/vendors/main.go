package vendors

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/chancegraff/project-news/pkg/services/collector/manager"
)

// Server will routinely push new articles into the database
type Server interface {
	Start(parent context.Context) error
	Run(parent context.Context) error
	Stop(parent context.Context) error
}

type server struct {
	manager *manager.Manager
}

// Start will run the server and block until it returns
func (s *server) Start(parent context.Context) error {
	_, cancel := context.WithCancel(parent)
	log.Println("Vendors server started")
	err := s.Run(parent)
	cancel()
	return err
}

// Run will periodically collect articles into the database until the context closes
func (s *server) Run(parent context.Context) error {
	for {
		select {
		case <-parent.Done():
			return errors.New("Vendors server shutting down")
		default:
			log.Println("Vendors server updated articles")
			time.Sleep(5 * time.Minute)
			s.Collect(parent)
		}
	}
}

func (s *server) Stop(parent context.Context) error {
	select {
	case <-parent.Done():
		log.Printf("Vendors stopped")
		return parent.Err()
	}
}

// NewServer ...
func NewServer(manager *manager.Manager) Server {
	return &server{
		manager: manager,
	}
}
