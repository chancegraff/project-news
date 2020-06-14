package main

import (
	"context"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/manager"
	"github.com/chancegraff/project-news/pkg/services/ranker/middlewares"
	"github.com/chancegraff/project-news/pkg/services/ranker/server"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// TODO Refactor calls to database into Manager
// TODO Create protocol buffer; compile into Golang; setup to be used with Collector service

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
	svc = middlewares.BindService(svc)
	endpoints := endpoints.NewEndpoints(svc)

	// Create the server
	server := server.NewHTTPServer(endpoints)
	defer server.Stop(ctx)

	// Start server
	go server.Start(ctx)

	// Bind until exit
	<-*done
	cancel()
}
