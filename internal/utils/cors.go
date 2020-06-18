package utils

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSOrigin ...
var CORSOrigin = handlers.AllowedOrigins(
	[]string{"*"},
)

// CORSHeaders ...
var CORSHeaders = handlers.AllowedHeaders(
	[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
)

// CORSMethods ...
var CORSMethods = handlers.AllowedMethods(
	[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
)

// CORSPolicy ...
var CORSPolicy = handlers.CORS(CORSMethods, CORSHeaders, CORSOrigin)

// SetCORSHeaders ...
func SetCORSHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, HEAD")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, X-Token-Auth, Authorization")
}
