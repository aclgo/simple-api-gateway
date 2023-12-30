package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		fmt.Fprint(w, "Internal Server Error")
	}
}
