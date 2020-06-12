package ranker

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// Start the service
	svc, err := newService()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create the endpoint handlers
	articlesHandler := httptransport.NewServer(
		makeArticlesEndpoint(svc),
		decodeArticlesRequest,
		encodeResponse,
	)

	userHandler := httptransport.NewServer(
		makeUserEndpoint(svc),
		decodeUserRequest,
		encodeResponse,
	)

	voteHandler := httptransport.NewServer(
		makeVoteEndpoint(svc),
		decodeVoteRequest,
		encodeResponse,
	)

	// Assign routes
	http.Handle("/articles", articlesHandler)
	http.Handle("/user", userHandler)
	http.Handle("/vote", voteHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":7998", nil))
	// Available at http://ranker.project-news-voter.app.localspace:7998/
}
