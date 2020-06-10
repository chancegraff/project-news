package collector

import (
	"log"
	"time"

	"github.com/chancegraff/project-news/internal/db"
	"github.com/chancegraff/project-news/internal/vendors"
	rest "github.com/chancegraff/project-news/pkg/collector"
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB
var arts *[]models.Article

func getArticles() {
	for {
		log.Println("Getting")
		arts = vendors.Get()
		store = db.FillArticles(store, arts)
		time.Sleep(5 * time.Minute)
	}
}

func StartServer() {
	log.Println("Running")

	arts = vendors.Get()
	store = db.Init()
	defer store.Close()

	go getArticles()

	rest.Listen(store)
}
