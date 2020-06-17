package auth

import (
	"context"
	"fmt"
	"time"

	pba "github.com/chancegraff/project-news/api/auth"
	"github.com/chancegraff/project-news/internal/utils"
	"google.golang.org/grpc"
)

// Proxy ...
type Proxy interface {
	Deregister(userID string) (*pba.User, error)
	Register(email string, password string) (*pba.User, error)
	User(userID string) (*pba.User, error)
	Verify(email string, password string) (*pba.User, error)
	Start(ctx context.Context) error
	Stop() error
}

type proxy struct {
	Address    string
	Connection *grpc.ClientConn
	Client     pba.AuthServiceClient
}

// NewProxy ...
func NewProxy() Proxy {
	port := utils.GetAuthPort()
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
func Connect(connection *grpc.ClientConn) pba.AuthServiceClient {
	return pba.NewAuthServiceClient(connection)
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
