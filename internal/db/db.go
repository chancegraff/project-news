package db

import (
	"github.com/chancegraff/project-news/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite adapter
)

// Init ...
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Vote{})
	db.AutoMigrate(&models.Client{})
	return db
}

// FillArticles puts articles in the DB
func FillArticles(db *gorm.DB, arts *[]models.Article) *gorm.DB {
	for _, a := range *arts {
		db.Where(models.Article{URL: a.URL}).FirstOrCreate(&a)
	}
	return db
}
