package token

import (
	"log"

	"github.com/chancegraff/project-news/internal/db"
	rest "github.com/chancegraff/project-news/pkg/token"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

func StartServer() {
	log.Println("Running")

	store = db.Init()
	defer store.Close()

	rest.Listen(store)
}
