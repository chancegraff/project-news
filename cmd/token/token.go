package main

import (
	"log"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/pkg/token"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

func TokenServer() {
	log.Println("Running")

	store = db.Init()
	defer store.Close()

	token.Listen(store)
}
