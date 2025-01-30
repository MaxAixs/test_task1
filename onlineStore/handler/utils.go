package handler

import (
	"encoding/json"
	"net/http"
)

func CheckContentType(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return false
	}

	return true
}

func ParseJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return false
	}

	return true
}
