package auth

import (
	"context"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/auth/endpoints"
	"github.com/chancegraff/project-news/pkg/services/auth/manager"
	"github.com/chancegraff/project-news/pkg/services/auth/middlewares"
	"github.com/chancegraff/project-news/pkg/services/auth/server"
	"github.com/chancegraff/project-news/pkg/services/auth/service"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// Run locally at 7997 and on the server at http://api.project-news-voter.app.localspace:7997/
func Run() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create database manager and service logger
	mgr := manager.NewManager()
	lgr := utils.Logger("auth")

	// Setup service and bind middlewares
	svc := service.NewService(&mgr)
	svc = middlewares.BindService(lgr, svc)

	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create RPC server
	srv := server.NewRPCServer(end)
	defer srv.Stop(ctx, lgr)

	// Start server and bind until exit
	go srv.Start(ctx, lgr)
	<-*done
	cancel()
}
