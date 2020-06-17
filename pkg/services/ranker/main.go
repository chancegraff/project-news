package ranker

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

// Run locally at 7998 and on the server at http://api.project-news-voter.app.localspace:7998/
func Run() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create database manager and service logger
	mgr := manager.NewManager()
	lgr := utils.Logger("ranker")

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
