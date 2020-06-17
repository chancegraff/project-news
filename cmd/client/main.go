package main

import (
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/gateway"
)

func main() {
	done := utils.GetDoneChannel()
	go gateway.Run()
	<-*done
}
