package router

import (
	"net/http"
	"taskapp/internal/handlers"
	"taskapp/internal/storage"

	"github.com/gorilla/mux"
)

// Setup configures all routes and returns the router
func Setup(store storage.Store) http.Handler {
	router := mux.NewRouter()
	taskHandler := handlers.NewTaskHandler(store)

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	api.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	api.HandleFunc("/health", taskHandler.HealthCheck).Methods("GET")

	// Enable CORS
	return enableCORS(router)
}

// enableCORS wraps the router with CORS headers
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
