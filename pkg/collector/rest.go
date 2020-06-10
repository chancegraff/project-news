package collector

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

// Listen will expose a port and watch it for requests
func Listen(api *mux.Router, s *gorm.DB) *mux.Router {
	store = s
	route := api.PathPrefix("/api/v1/articles").Subrouter()

	route.HandleFunc("", all).Methods("GET", "OPTIONS")
	route.HandleFunc("", all).Methods("GET", "OPTIONS").Queries("offset", "{offset:[0-9]+}")
	route.HandleFunc("/get", get).Methods("GET", "OPTIONS").Queries("id", "{id:[0-9]+}")

	return route
}
