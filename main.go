package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/internal/vendors"
	"github.com/chancegraff/project-news/pkg/auth"
	"github.com/chancegraff/project-news/pkg/collector"
	"github.com/chancegraff/project-news/pkg/ranker"
	"github.com/chancegraff/project-news/pkg/token"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getArticles(store *gorm.DB) {
	for {
		db.FillArticles(store, vendors.Get())
		log.Println("Server got articles")
		time.Sleep(5 * time.Minute)
	}
}

func main() {
	port := ":" + utils.GetEnv("PORT", "3000")
	path, _ := os.Getwd()
	fp := filepath.Join(path, "web", "build")

	store := db.Init()
	defer store.Close()

	fs := http.FileServer(http.Dir(fp))
	rt := mux.NewRouter()

	rt.PathPrefix("/").Handler(fs)

	api := rt.PathPrefix("/api/v1").Subrouter()

	api = auth.Listen(api, store)
	api = collector.Listen(api, store)
	api = ranker.Listen(api, store)
	api = token.Listen(api, store)

	go getArticles(store)

	log.Println("Server is running")

	log.Fatal(
		http.ListenAndServe(
			port,
			handlers.CORS(
				handlers.AllowedHeaders(
					[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
				),
				handlers.AllowedMethods(
					[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
				),
				handlers.AllowedOrigins(
					[]string{"http://localhost:3000"},
				),
			)(rt),
		),
	)
}
