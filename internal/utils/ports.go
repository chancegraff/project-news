package utils

import (
	"os"
	"strconv"
)

// GetCollectorPort will return the collector port value
func GetCollectorPort() int {
	collectorPortString := os.Getenv("COLLECTOR_PORT")
	collectorPort, err := strconv.Atoi(collectorPortString)
	if err != nil {
		collectorPort = 7999
	}
	return collectorPort
}

// GetRankerPort will return the ranker port value
func GetRankerPort() int {
	rankerPortString := os.Getenv("RANKER_PORT")
	rankerPort, err := strconv.Atoi(rankerPortString)
	if err != nil {
		rankerPort = 7998
	}
	return rankerPort
}
