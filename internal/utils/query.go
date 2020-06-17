package utils

import (
	"os"
	"strconv"
)

// GetQueryRateLimit will return the query rate limit per second
func GetQueryRateLimit() int {
	queryRateString := os.Getenv("QUERIES_PER_SECOND")
	queryRate, err := strconv.Atoi(queryRateString)
	if err != nil {
		queryRate = 100
	}
	return queryRate
}
