package ranker

import (
	"log"

	"github.com/chancegraff/project-news/pkg/services/ranker/endpoints"
	"github.com/chancegraff/project-news/pkg/services/ranker/server"
	"github.com/chancegraff/project-news/pkg/services/ranker/service"
)

func main() {
	// Setup the service
	svc := service.NewService()
	endpoints := endpoints.NewEndpoints(svc)
	server := server.NewHTTPServer(endpoints)

	// Start server at http://api.project-news-voter.app.localspace:7998/
	log.Fatal(server.Start(":7998"))
}
