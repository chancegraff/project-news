package collector

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var store *gorm.DB

// Listen will expose a port and watch it for requests
func Listen(s *gorm.DB) {
	store = s
	rt := mux.NewRouter()
	api := rt.PathPrefix("/api/v1/articles").Subrouter()

	api.HandleFunc("", all).Methods("GET", "OPTIONS")
	api.HandleFunc("", all).Methods("GET", "OPTIONS").Queries("offset", "{offset:[0-9]+}")
	api.HandleFunc("/get", get).Methods("GET", "OPTIONS").Queries("id", "{id:[0-9]+}")

	log.Println("Listening")

	log.Fatal(http.ListenAndServe(
		":8000",
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
