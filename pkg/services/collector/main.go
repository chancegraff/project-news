package collector

import (
	"context"
	"os"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/manager"
	"github.com/chancegraff/project-news/pkg/services/collector/middlewares"
	"github.com/chancegraff/project-news/pkg/services/collector/server"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	"github.com/chancegraff/project-news/pkg/services/collector/vendors"
	"github.com/go-kit/kit/log"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

// Run locally at 7999 and on the server at http://api.project-news-voter.app.localspace:7999/
func Run() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create database manager
	mgr := manager.NewManager()

	// Setup the endpoints
	svc := service.NewService(ctx, &mgr)

	// Create logger
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "collector")

	// Bind middleware
	svc = middlewares.BindService(logger, svc)

	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create Collector server
	clctr := vendors.NewServer(&mgr)
	defer clctr.Stop(ctx)

	// Create RPC server
	srv := server.NewRPCServer(end)
	defer srv.Stop(ctx, logger)

	// Start servers
	go clctr.Start(ctx)
	go srv.Start(ctx, logger)

	// Bind until exit
	<-*done
	cancel()
}
