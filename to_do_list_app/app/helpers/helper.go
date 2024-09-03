package helpers

import (
	"encoding/json"
	"net/http"
)

// ParseRequest decodes the JSON body of the HTTP request into provided payload.
func ParseRequest(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

// SendResponse sends a JSON response with the given status code and payload.
func SendResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var response any

	if msg, ok := payload.(string); ok {
		response = map[string]string{"message": msg}
	} else {
		response = payload
	}

	json.NewEncoder(w).Encode(response)
}
