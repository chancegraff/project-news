package collector

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/vendors"
	"github.com/chancegraff/project-news/pkg/collector"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload" // Autoloading .env
)

func fillArticles(parent context.Context, store *gorm.DB) {
	for {
		select {
		case <-parent.Done():
			return
		default:
			db.FillArticles(store, vendors.Get())
			log.Println("Server got articles")
		}
		time.Sleep(5 * time.Minute)
	}
}

// StartServer will start the Collector process
func StartServer() {
	log.Println("Running")

	// Open the DB
	store := db.Init()
	defer store.Close()

	// Bind to context
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	defer signal.Stop(done)

	// Call routines
	go fillArticles(ctx, store)
	go collector.Listen(ctx, store)

	// Block until interrupted
	log.Println("Server is running")
	<-done
	cancel()

	// Close server gracefully
	log.Println("Server is closing")
	os.Exit(0)
}
