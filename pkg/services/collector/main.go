package collector

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

// Run locally at 7999 and on the server at http://api.project-news-voter.app.localspace:7999/
func Run() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create database manager and service logger
	mgr := manager.NewManager()
	lgr := utils.Logger("collector")

	// Setup service and bind middlewares
	svc := service.NewService(ctx, &mgr)
	svc = middlewares.BindService(lgr, svc)

	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create Collector server
	clctr := vendors.NewServer(&mgr)
	defer clctr.Stop(ctx, lgr)

	// Create RPC server
	srv := server.NewRPCServer(end)
	defer srv.Stop(ctx, lgr)

	// Start servers and bind until exit
	go clctr.Start(ctx, lgr)
	go srv.Start(ctx, lgr)
	<-*done
	cancel()
}
