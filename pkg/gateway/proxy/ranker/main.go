package ranker

import (
	"context"
	"time"

	pbr "github.com/chancegraff/project-news/api/ranker"
	"github.com/chancegraff/project-news/internal/utils"
	"google.golang.org/grpc"
)

// Proxy ...
type Proxy interface {
	Articles(articleIDs []string) ([]*pbr.ArticleVotes, error)
	User(userID string) (*pbr.UserVotes, error)
	Vote(articleID, userID string) (*pbr.ArticleVotes, error)
	Start(ctx context.Context) error
	Stop() error
}

type proxy struct {
	Address    string
	Connection *grpc.ClientConn
	Client     pbr.RankerServiceClient
}

// NewProxy ...
func NewProxy() Proxy {
	port := utils.GetRankerPort()
	address := utils.GetAPIAddress(port)
	return &proxy{
		Address: address,
	}
}

// Dial ...
func Dial(ctx context.Context, address string) (*grpc.ClientConn, error) {
	connection, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// Connect ...
func Connect(connection *grpc.ClientConn) pbr.RankerServiceClient {
	return pbr.NewRankerServiceClient(connection)
}

// Start ...
func (p *proxy) Start(ctx context.Context) error {
	connection, err := Dial(ctx, p.Address)
	if err != nil {
		return err
	}
	p.Connection = connection
	p.Client = Connect(p.Connection)
	return nil
}

// Stop ...
func (p *proxy) Stop() error {
	return p.Connection.Close()
}
