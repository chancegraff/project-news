package main

import (
	"log"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/pkg/auth"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

func AuthServer() {
	log.Println("Running")

	store = db.Init()
	defer store.Close()

	auth.Listen(store)
}
