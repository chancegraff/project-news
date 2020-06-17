package service

import (
	"context"
	"fmt"
	"time"

	"github.com/chancegraff/project-news/internal/models"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/manager"
	"google.golang.org/grpc"

	pb "github.com/chancegraff/project-news/api/ranker"
)

// Service implements the collector interface
type Service interface {
	All(offset int) ([]models.Article, error)
	Get(id int) (models.Article, error)
}

type service struct {
	Manager                 *manager.Manager
	RankerServiceAddress    string
	RankerServiceConnection *grpc.ClientConn
	RankerService           pb.RankerServiceClient
}

// NewService instantiates the service with a connection to the database
func NewService(ctx context.Context, manager *manager.Manager) Service {
	port := utils.GetRankerPort()
	address := fmt.Sprint(":", port)
	rankerConnection, err := grpc.DialContext(ctx, address,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3))
	if err != nil {
		panic(err)
	}
	return &service{
		Manager:                 manager,
		RankerServiceConnection: rankerConnection,
		RankerService:           pb.NewRankerServiceClient(rankerConnection),
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Service) Service
