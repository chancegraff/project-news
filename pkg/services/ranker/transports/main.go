package transports

import (
	"context"
	"encoding/json"
	"net/http"
)

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
