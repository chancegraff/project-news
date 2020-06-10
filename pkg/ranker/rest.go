package ranker

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

// Listen ...
func Listen(api *mux.Router, s *gorm.DB) *mux.Router {
	store = s
	route := api.PathPrefix("/ranks").Subrouter()

	route.HandleFunc("/articles", articles).Methods("POST", "OPTIONS")
	route.HandleFunc("/user", user).Methods("POST", "OPTIONS")
	route.HandleFunc("/vote", vote).Methods("POST", "OPTIONS")

	return api
}
