package storage

import "taskapp/internal/models"

// Store defines the interface for task storage
type Store interface {
	GetAll() []*models.Task
	GetByID(id string) (*models.Task, bool)
	Save(task *models.Task)
	Delete(id string) bool
	GetNextID() int
}
