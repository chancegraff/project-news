package main

import (
	"log"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/pkg/ranker"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

func RankerServer() {
	log.Println("Running")

	store = db.Init()
	defer store.Close()

	ranker.Listen(store)
}
