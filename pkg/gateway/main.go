package main

import (
	"context"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway/endpoints"
	"github.com/chancegraff/project-news/pkg/gateway/middlewares"
	"github.com/chancegraff/project-news/pkg/gateway/server"
	"github.com/chancegraff/project-news/pkg/gateway/service"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Bind resources and create logger
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()
	lgr := utils.Logger("gateway")

	// Create service with middlewares and endpoints
	svc := service.NewService(ctx)
	svc = middlewares.BindService(lgr, svc)
	end := endpoints.NewEndpoints(svc)

	// Create server
	srv := server.NewServer(end, lgr)
	defer srv.Stop(ctx)

	// Start server and bind until exit
	go srv.Start(ctx)
	<-*done
	cancel()
}
