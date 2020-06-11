package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/auth"
	"github.com/chancegraff/project-news/pkg/collector"
	"github.com/chancegraff/project-news/pkg/ranker"
	"github.com/chancegraff/project-news/pkg/token"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // Autoloading .env
)

var getCORS = handlers.CORS(
	handlers.AllowedHeaders(
		[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
	),
	handlers.AllowedMethods(
		[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	),
)

func main() {
	log.Println("Server is starting")

	store := db.Init()
	defer store.Close()

	rt := mux.NewRouter()

	auth.Listen(rt, store)
	collector.Listen(rt, store)
	ranker.Listen(rt, store)
	token.Listen(rt, store)

	path, _ := os.Getwd()
	fp := filepath.Join(path, "web", "build")
	fs := http.FileServer(http.Dir(fp))
	rt.PathPrefix("/").Handler(fs)

	var wait time.Duration
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	port := fmt.Sprintf(":%s", utils.GetEnv("PORT", "3000"))
	srv := &http.Server{
		Handler:      getCORS(rt),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	defer srv.Shutdown(ctx)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Server is running")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Server is closing")

	os.Exit(0)
}
