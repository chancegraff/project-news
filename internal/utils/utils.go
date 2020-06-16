package utils

import (
	"os"
	"os/signal"
	"time"
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
