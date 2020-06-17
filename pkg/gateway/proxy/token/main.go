package token

import (
	"context"
	"fmt"
	"time"

	pbt "github.com/chancegraff/project-news/api/token"
	"github.com/chancegraff/project-news/internal/utils"
	"google.golang.org/grpc"
)

// Proxy ...
type Proxy interface {
	Generate(identifiers *pbt.Identifiers, client *pbt.Client) (string, error)
	Verify(identifiers *pbt.Identifiers, client *pbt.Client) (string, error)
	Start(ctx context.Context) error
	Stop() error
}

type proxy struct {
	Address    string
	Connection *grpc.ClientConn
	Client     pbt.TokenServiceClient
}

// NewProxy ...
func NewProxy() Proxy {
	port := utils.GetTokenPort()
	address := fmt.Sprint(":", port)
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
func Connect(connection *grpc.ClientConn) pbt.TokenServiceClient {
	return pbt.NewTokenServiceClient(connection)
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
