package utils

import "os"

// GetRankerAddress returns the address for the ranker API
func GetRankerAddress() string {
	address := os.Getenv("RANKER_HTTP_ADDRESS")
	if address == "" {
		address = "http://ranker.project-news-voter.app.localspace:7998/"
	}
	return address
}
