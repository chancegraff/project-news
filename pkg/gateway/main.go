package main

import (
	"context"
	"os"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/server"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	"github.com/go-kit/kit/log"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create service
	svc := service.NewService(ctx)

	// Create middleware
	// mdl, lgr := middlewares.NewMiddleware(svc)

	// Bind service to middleware
	// svc = mdl.Bind()
	lgr := log.NewLogfmtLogger(os.Stderr)
	lgr = log.With(lgr, "ts", log.DefaultTimestampUTC)
	lgr = log.With(lgr, "caller", log.DefaultCaller)
	lgr = log.With(lgr, "service", "gateway")
	// Create endpoints
	end := endpoints.NewEndpoints(svc)

	// Create server
	srv := server.NewServer(end, lgr)
	defer srv.Stop(ctx)

	// Start server
	go srv.Start(ctx)

	// Bind until exit
	<-*done
	cancel()
}
