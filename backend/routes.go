package main

import (
	"github.com/gorilla/mux"
)

// setupRoutes configures all API endpoints
func setupRoutes(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	
	// Task endpoints
	api.HandleFunc("/tasks", getTasks).Methods("GET")
	api.HandleFunc("/tasks", createTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	
	// Health check endpoint
	api.HandleFunc("/health", healthCheck).Methods("GET")
}
