package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func ValidateSeller(name, phone string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if phone == "" {
		return fmt.Errorf("phone is required")
	}

	return nil
}

func getSellerID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		return 0, fmt.Errorf("seller ID is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid seller ID")
	}

	return id, nil
}
