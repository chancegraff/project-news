package utils

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
)

// GetDoneChannel creates a new channel to listen for done signals on
func GetDoneChannel() *chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	return &done
}

// Tomorrow will return a time.Time instance for tomorrow
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// Logger will return a gokit logger with default params
func Logger(serviceName string) log.Logger {
	lgr := log.NewLogfmtLogger(os.Stderr)
	lgr = log.WithPrefix(lgr, "ts", log.DefaultTimestampUTC)
	lgr = log.WithPrefix(lgr, "service", serviceName)
	return lgr
}

// Getwd ...
func Getwd() string {
	wd := os.Getenv("WORKING_DIRECTORY")
	if wd == "" {
		wd = "/Users/chancegraff/Repositories/project-news"
	}
	return wd
}

// GetAPIAddress ...
func GetAPIAddress(p int) string {
	port := strconv.Itoa(p)
	host := os.Getenv("API_HOST")
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

// GetEnvironment ...
func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}
	return env
}
