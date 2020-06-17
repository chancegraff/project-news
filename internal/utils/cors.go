package utils

import "github.com/gorilla/handlers"

// CORSHeaders ...
var CORSHeaders = handlers.AllowedHeaders(
	[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
)

// CORSMethods ...
var CORSMethods = handlers.AllowedMethods(
	[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
)

// CORSPolicy ...
var CORSPolicy = handlers.CORS(CORSMethods, CORSHeaders)
