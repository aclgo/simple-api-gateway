package auth

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error   string
	Message string
}

func Json(w http.ResponseWriter, v any, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
