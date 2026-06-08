package main

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}
