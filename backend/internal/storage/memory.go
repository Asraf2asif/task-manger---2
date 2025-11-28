package storage

import (
	_ "embed"
	"encoding/json"
	"log"
	"sync"
	"taskapp/internal/models"
	"time"
)

//go:embed demo_tasks.json
var demoTasksJSON []byte

// MemoryStore handles in-memory task storage
type MemoryStore struct {
	tasks  map[string]*models.Task
	mutex  *sync.RWMutex
	nextID int
}

// NewMemoryStore creates a new in-memory storage instance
func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		tasks:  make(map[string]*models.Task),
		mutex:  &sync.RWMutex{},
		nextID: 1,
	}
	store.seedDemoTasks()
	return store
}

// DemoTask represents the JSON structure for demo tasks
type DemoTask struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

// seedDemoTasks loads demo data from JSON file
func (s *MemoryStore) seedDemoTasks() {
	var demoTasks []DemoTask

	if err := json.Unmarshal(demoTasksJSON, &demoTasks); err != nil {
		log.Printf("Error loading demo tasks: %v", err)
		return
	}

	now := time.Now()
	for i, dt := range demoTasks {
		// Create timestamps with varying ages
		hoursAgo := time.Duration((i+1)*4) * time.Hour
		task := &models.Task{
			ID:          dt.ID,
			Title:       dt.Title,
			Description: dt.Description,
			Status:      dt.Status,
			Priority:    dt.Priority,
			CreatedAt:   now.Add(-hoursAgo),
			UpdatedAt:   now.Add(-hoursAgo / 2),
		}
		s.tasks[task.ID] = task
	}

	s.nextID = 73
	log.Printf("Loaded %d demo tasks from JSON", len(demoTasks))
}

// GetAll returns all tasks (thread-safe)
func (s *MemoryStore) GetAll() []*models.Task {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]*models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result
}

// GetByID retrieves a single task (thread-safe)
func (s *MemoryStore) GetByID(id string) (*models.Task, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

// Save stores or updates a task (thread-safe)
func (s *MemoryStore) Save(task *models.Task) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.tasks[task.ID] = task
}

// Delete removes a task (thread-safe)
func (s *MemoryStore) Delete(id string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.tasks[id]; exists {
		delete(s.tasks, id)
		return true
	}
	return false
}

// GetNextID returns the next available ID
func (s *MemoryStore) GetNextID() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.nextID
	s.nextID++
	return id
}
