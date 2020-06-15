package main

import (
	"context"
	"os"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/auth/endpoints"
	"github.com/chancegraff/project-news/pkg/services/auth/manager"
	"github.com/chancegraff/project-news/pkg/services/auth/middlewares"
	"github.com/chancegraff/project-news/pkg/services/auth/server"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	"github.com/go-kit/kit/log"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// Runs locally at 7997 and on the server at:
// http://api.project-news-voter.app.localspace:7997/
func main() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create database manager
	mgr := manager.NewManager()

	// Setup the endpoints
	svc := service.NewService(&mgr)

	// Create logger
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "auth")

	// Bind middleware
	svc = middlewares.BindService(logger, svc)

	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create RPC server
	srv := server.NewRPCServer(end)
	defer srv.Stop(ctx, logger)

	// Start servers
	go srv.Start(ctx, logger)

	// Bind until exit
	<-*done
	cancel()
}
