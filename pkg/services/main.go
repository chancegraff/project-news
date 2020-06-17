package main

import (
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services/auth"
	"github.com/chancegraff/project-news/pkg/services/collector"
	"github.com/chancegraff/project-news/pkg/services/ranker"
	"github.com/chancegraff/project-news/pkg/services/token"
)

func main() {
	done := utils.GetDoneChannel()

	go auth.Run()
	go collector.Run()
	go ranker.Run()
	go token.Run()

	<-*done
}
