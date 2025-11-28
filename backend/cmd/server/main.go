package main

import (
	"log"
	"net/http"
	"os"
	"taskapp/internal/router"
	"taskapp/internal/storage"
)

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize storage
	store := storage.NewMemoryStore()

	// Setup router
	handler := router.Setup(store)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
