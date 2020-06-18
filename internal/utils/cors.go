package utils

import "github.com/gorilla/handlers"

// CORSOrigin ...
var CORSOrigin = handlers.AllowedOrigins(
	[]string{"*"},
)

// CORSHeaders ...
var CORSHeaders = handlers.AllowedHeaders(
	[]string{"*"},
)

// CORSMethods ...
var CORSMethods = handlers.AllowedMethods(
	[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
)

// CORSPolicy ...
var CORSPolicy = handlers.CORS(CORSMethods, CORSHeaders, CORSOrigin)
