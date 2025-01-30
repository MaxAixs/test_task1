package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func newSuccessResponse(w http.ResponseWriter, data interface{}) {
	log.Printf("status %d", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"data": data}); err != nil {
		log.Printf("failed to write JSON response: %v", err)
	}
}
