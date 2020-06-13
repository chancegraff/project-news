package collector

import (
	"log"

	"github.com/chancegraff/project-news/pkg/services/collector/endpoints"
	"github.com/chancegraff/project-news/pkg/services/collector/middlewares"
	"github.com/chancegraff/project-news/pkg/services/collector/server"
	"github.com/chancegraff/project-news/pkg/services/collector/service"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Setup the service
	svc := service.NewService()
	svc = middlewares.BindService(svc)
	endpoints := endpoints.NewEndpoints(svc)
	server := server.NewHTTPServer(endpoints)

	// Start server at http://api.project-news-voter.app.localspace:7999/
	log.Println("Server started at :7998")
	log.Fatal(server.Start(":7998"))
}
