package ranker

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

// Listen ...
func Listen(s *gorm.DB) {
	store = s
	rt := mux.NewRouter()
	api := rt.PathPrefix("/api/v1/ranks").Subrouter()

	api.HandleFunc("/articles", articles).Methods("POST", "OPTIONS")
	api.HandleFunc("/user", user).Methods("POST", "OPTIONS")
	api.HandleFunc("/vote", vote).Methods("POST", "OPTIONS")

	log.Println("Listening")
	log.Fatal(http.ListenAndServe(
		":8001",
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
			),
			handlers.AllowedMethods(
				[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
			),
			handlers.AllowedOrigins(
				[]string{"http://localhost:3000"},
			),
		)(rt),
	))
}
