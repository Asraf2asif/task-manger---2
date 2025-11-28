package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// getTasks returns all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := getAllTasks()
	respondJSON(w, http.StatusOK, tasks)
}

// createTask adds a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	// Validate required fields
	if task.Title == "" {
		respondError(w, http.StatusBadRequest, "Title is required", nil)
		return
	}
	
	// Set defaults
	task.ID = fmt.Sprintf("%d", nextID)
	nextID++
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	
	if task.Status == "" {
		task.Status = "todo"
	}
	if task.Priority == "" {
		task.Priority = "medium"
	}
	
	saveTask(&task)
	respondJSON(w, http.StatusCreated, task)
}

// updateTask modifies an existing task
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	existingTask, exists := getTaskByID(id)
	if !exists {
		respondError(w, http.StatusNotFound, "Task not found", nil)
		return
	}
	
	var updates Task
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	// Update fields while preserving ID and CreatedAt
	existingTask.Title = updates.Title
	existingTask.Description = updates.Description
	existingTask.Status = updates.Status
	existingTask.Priority = updates.Priority
	existingTask.UpdatedAt = time.Now()
	
	saveTask(existingTask)
	respondJSON(w, http.StatusOK, existingTask)
}

// deleteTask removes a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if !removeTask(id) {
		respondError(w, http.StatusNotFound, "Task not found", nil)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

// healthCheck returns server status
func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}
