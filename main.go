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
)

func main() {
	go auth.StartServer()
	go collector.StartServer()
	go ranker.StartServer()
	go token.StartServer()

	path, _ := os.Getwd()
	fp := filepath.Join(path, "web", "build")
	fs := http.FileServer(http.Dir(fp))

	http.Handle("/", fs)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
