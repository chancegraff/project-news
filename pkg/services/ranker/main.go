package main

import (
	"context"
	"os"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/manager"
	"github.com/chancegraff/project-news/pkg/services/ranker/middlewares"
	"github.com/chancegraff/project-news/pkg/services/ranker/server"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	"github.com/go-kit/kit/log"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// Runs locally at 7998 and on the server at:
// http://api.project-news-voter.app.localspace:7998/
func main() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Connect database manager
	mgr := manager.NewManager()

	// Setup the service
	svc := service.NewService(&mgr)

	// Create logger
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// Bind middleware
	svc = middlewares.BindService(logger, svc)

	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create RPC server
	srv := server.NewRPCServer(end)
	defer srv.Stop(ctx, logger)

	// Start server
	go srv.Start(ctx, logger)

	// // Create the server
	// srv := server.NewHTTPServer(endpoints)
	// defer srv.Stop(ctx, logger)

	// // Start server
	// go srv.Start(ctx, logger)

	// Bind until exit
	<-*done
	cancel()
}
