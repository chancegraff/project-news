package collector

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service{}

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

	http.Handle("/get", getHandler)
	http.Handle("/all", allHandler)

	// Available at http://collector.project-news-voter.app.localspace:7999/
	log.Fatal(http.ListenAndServe(":7999", nil))
}
