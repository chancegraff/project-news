package main

import (
	"context"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/manager"
	"github.com/chancegraff/project-news/pkg/services/collector/middlewares"
	"github.com/chancegraff/project-news/pkg/services/collector/server"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/vendors"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// Runs localy at 7999 and on the server at:
// http://api.project-news-voter.app.localspace:7999/
func main() {
	// Create database manager
	mgr := manager.NewManager()

	// Setup the endpoints
	svc := service.NewService(&mgr)
	svc = middlewares.BindService(svc)
	endpoints := endpoints.NewEndpoints(svc)

	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create HTTP server
	server := server.NewHTTPServer(endpoints)
	defer server.Stop(ctx)

	// Create Collector server
	collector := vendors.NewServer(&mgr)
	defer collector.Stop(ctx)

	// Start servers
	go collector.Start(ctx)
	go server.Start(ctx)

	// Bind until exit
	<-*done
	cancel()
}
