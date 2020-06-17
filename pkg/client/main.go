package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chancegraff/project-news/internal/utils"
)

// Run will handle the lifecycle of the client server
func Run() {
	log.Println("Client server starting")

	// Bind resources
	done := utils.GetDoneChannel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Build file server and router
	path, _ := os.Getwd()
	fp := filepath.Join(path, "web", "build")
	fs := http.FileServer(http.Dir(fp))
	rt := http.NewServeMux()
	rt.Handle("/", fs)

	// Build http server
	port := utils.GetClientPort()
	address := fmt.Sprintf(":%v", port)
	srv := &http.Server{
		Handler:      utils.CORSPolicy(rt),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	defer srv.Shutdown(ctx)

	// Run server
	go func() {
		log.Printf("Client server running at %s", address)
		if err := srv.ListenAndServe(); err != nil {
			cancel()
			*done <- os.Interrupt
		}
	}()

	// Bind until exit
	<-*done
	cancel()
}
