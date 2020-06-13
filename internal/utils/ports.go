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
		collectorPort = 7998
	}
	return collectorPort
}
