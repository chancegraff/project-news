package utils

import (
	"os"
	"os/signal"
)

// GetDoneChannel creates a new channel to listen for done signals on
func GetDoneChannel() *chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	return &done
}
