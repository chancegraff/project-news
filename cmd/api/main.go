package main

import (
	"github.com/chancegraff/project-news/internal/utils"
	"github.com/chancegraff/project-news/pkg/services"
)

func main() {
	done := utils.GetDoneChannel()
	go services.Run()
	<-*done
}
