package main

import (
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/client"
)

func main() {
	done := utils.GetDoneChannel()
	go client.Run()
	<-*done
}
