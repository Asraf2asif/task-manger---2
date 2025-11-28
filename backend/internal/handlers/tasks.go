package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taskapp/internal/models"
	"taskapp/internal/storage"
	"time"

	"github.com/gorilla/mux"
)

// TaskHandler handles task-related HTTP requests
type TaskHandler struct {
	store storage.Store
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(store storage.Store) *TaskHandler {
	return &TaskHandler{store: store}
}

// GetTasks returns all tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.store.GetAll()
	respondJSON(w, http.StatusOK, tasks)
}

// CreateTask adds a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

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
	task.ID = fmt.Sprintf("%d", h.store.GetNextID())
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if task.Status == "" {
		task.Status = "todo"
	}
	if task.Priority == "" {
		task.Priority = "medium"
	}

	h.store.Save(&task)
	respondJSON(w, http.StatusCreated, task)
}

// UpdateTask modifies an existing task
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	existingTask, exists := h.store.GetByID(id)
	if !exists {
		respondError(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	var updates models.Task
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

	h.store.Save(existingTask)
	respondJSON(w, http.StatusOK, existingTask)
}

// DeleteTask removes a task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !h.store.Delete(id) {
		respondError(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HealthCheck returns server status
func (h *TaskHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}
