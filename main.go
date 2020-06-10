package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chancegraff/project-news/cmd/main/auth"
	"github.com/chancegraff/project-news/cmd/main/collector"
	"github.com/chancegraff/project-news/cmd/main/ranker"
	"github.com/chancegraff/project-news/cmd/main/token"
	"github.com/chancegraff/project-news/internal/utils"
)

func main() {
	go auth.StartServer()
	go collector.StartServer()
	go ranker.StartServer()
	go token.StartServer()

	port := utils.GetEnv("PORT", "3000")
	path, _ := os.Getwd()
	fp := filepath.Join(path, "web", "build")
	fs := http.FileServer(http.Dir(fp))

	http.Handle("/", fs)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
