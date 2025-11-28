package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"taskapp/internal/models"
)

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// respondError sends an error response
func respondError(w http.ResponseWriter, status int, message string, err error) {
	if err != nil {
		log.Printf("Error: %s - %v", message, err)
	}

	response := models.ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	}

	respondJSON(w, status, response)
}
