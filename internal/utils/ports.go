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

// GetAuthPort will return the ranker port value
func GetAuthPort() int {
	authPortString := os.Getenv("AUTH_PORT")
	authPort, err := strconv.Atoi(authPortString)
	if err != nil {
		authPort = 7997
	}
	return authPort
}

// GetTokenPort will return the ranker port value
func GetTokenPort() int {
	tokenPortString := os.Getenv("TOKEN_PORT")
	tokenPort, err := strconv.Atoi(tokenPortString)
	if err != nil {
		tokenPort = 7996
	}
	return tokenPort
}

// GetGatewayPort will return the gateway port value
func GetGatewayPort() int {
	gatewayPortString := os.Getenv("GATEWAY_PORT")
	gatewayPort, err := strconv.Atoi(gatewayPortString)
	if err != nil {
		gatewayPort = 8000
	}
	return gatewayPort
}

// GetClientPort will return the gateway port value
func GetClientPort() int {
	clientPortString := os.Getenv("CLIENT_PORT")
	clientPort, err := strconv.Atoi(clientPortString)
	if err != nil {
		clientPort = 8080
	}
	return clientPort
}
