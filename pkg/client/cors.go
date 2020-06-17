package client

import "github.com/gorilla/handlers"

// Headers ...
var Headers = handlers.AllowedHeaders(
	[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
)

// Methods ...
var Methods = handlers.AllowedMethods(
	[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
)

// Policy ...
var Policy = handlers.CORS(Headers, Methods)
