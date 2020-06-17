package main

import (
	"context"
	"log"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/middlewares"
	"github.com/chancegraff/project-news/pkg/gateway/proxy"
	"github.com/chancegraff/project-news/pkg/gateway/server"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Bind resources
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()

	// Create logger and middlewares
	lgr := utils.Logger("gateway")
	mdl := middlewares.NewMiddlewares(lgr)

	// Create service proxies
	prx, err := proxy.NewProxy(ctx, lgr)
	if err != nil {
		log.Fatal("Failed to create proxies", err)
		return
	}

	// Create service and bind to middleware
	svc := service.NewService(prx)
	svc = mdl.BindService(svc)

	// Create endpoints and bind to middleware
	end := endpoints.NewEndpoints(svc)
	end = mdl.BindEndpoints(end)

	// Create server
	srv := server.NewServer(end, lgr)
	defer srv.Stop(ctx)

	// Start server and bind until exit
	go srv.Start(ctx)
	<-*done
	cancel()
}
