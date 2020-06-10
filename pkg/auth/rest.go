package auth

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

// Listen ...
func Listen(api *mux.Router, s *gorm.DB) *mux.Router {
	store = s
	route := api.PathPrefix("/api/v1/auth").Subrouter()

	route.HandleFunc("/deregister", deregister).Methods("POST", "OPTIONS")
	route.HandleFunc("/register", register).Methods("POST", "OPTIONS")
	route.HandleFunc("/verify", verify).Methods("POST", "OPTIONS")
	route.HandleFunc("/user", user).Methods("POST", "OPTIONS")

	return api
}
