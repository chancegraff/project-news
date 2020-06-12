package collector

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// Start the service
	var svc Service
	svc, err := newService()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Setup middlewares
	svc = articlesEndpointMiddleware("http://ranker.project-news-voter.app.localspace:7998/")(svc)

	// Create the endpoint handlers
	getHandler := httptransport.NewServer(
		makeGetEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

	allHandler := httptransport.NewServer(
		makeAllEndpoint(svc),
		decodeAllRequest,
		encodeResponse,
	)

	// Assign routes
	http.Handle("/get", getHandler)
	http.Handle("/all", allHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":7999", nil))
	// Available at http://collector.project-news-voter.app.localspace:7999/
}
