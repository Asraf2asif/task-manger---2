package storage

import (
	"sync"
	"taskapp/internal/models"
	"time"
)

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

// seedDemoTasks creates initial demo data
func (s *MemoryStore) seedDemoTasks() {
	demoTasks := []*models.Task{
		{
			ID:          "1",
			Title:       "Review pull requests",
			Description: "Check the pending PRs from the team",
			Status:      "todo",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-48 * time.Hour),
		},
		{
			ID:          "2",
			Title:       "Update API documentation",
			Description: "Document the new endpoints we added last week",
			Status:      "in-progress",
			Priority:    "medium",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * time.Hour),
		},
		{
			ID:          "3",
			Title:       "Fix login bug",
			Description: "Users reported issues with OAuth login flow",
			Status:      "done",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-72 * time.Hour),
			UpdatedAt:   time.Now().Add(-12 * time.Hour),
		},
		{
			ID:          "4",
			Title:       "Implement dark mode",
			Description: "Add theme toggle for better user experience",
			Status:      "in-progress",
			Priority:    "low",
			CreatedAt:   time.Now().Add(-36 * time.Hour),
			UpdatedAt:   time.Now().Add(-10 * time.Hour),
		},
		{
			ID:          "5",
			Title:       "Optimize database queries",
			Description: "Some queries are taking too long, need indexing",
			Status:      "in-progress",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-20 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
		{
			ID:          "6",
			Title:       "Write unit tests",
			Description: "Increase test coverage to at least 80%",
			Status:      "todo",
			Priority:    "medium",
			CreatedAt:   time.Now().Add(-12 * time.Hour),
			UpdatedAt:   time.Now().Add(-12 * time.Hour),
		},
		{
			ID:          "7",
			Title:       "Setup CI/CD pipeline",
			Description: "Automate deployment process with GitHub Actions",
			Status:      "done",
			Priority:    "medium",
			CreatedAt:   time.Now().Add(-96 * time.Hour),
			UpdatedAt:   time.Now().Add(-24 * time.Hour),
		},
		{
			ID:          "8",
			Title:       "Refactor authentication module",
			Description: "Clean up the auth code and improve security",
			Status:      "in-progress",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-8 * time.Hour),
			UpdatedAt:   time.Now().Add(-3 * time.Hour),
		},
		{
			ID:          "9",
			Title:       "Design new landing page",
			Description: "Create mockups for the updated homepage",
			Status:      "in-progress",
			Priority:    "medium",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-6 * time.Hour),
		},
		{
			ID:          "10",
			Title:       "Mobile app research",
			Description: "Evaluate React Native vs Flutter for mobile version",
			Status:      "todo",
			Priority:    "low",
			CreatedAt:   time.Now().Add(-4 * time.Hour),
			UpdatedAt:   time.Now().Add(-4 * time.Hour),
		},
		{
			ID:          "11",
			Title:       "Performance monitoring setup",
			Description: "Integrate New Relic or DataDog for monitoring",
			Status:      "in-progress",
			Priority:    "medium",
			CreatedAt:   time.Now().Add(-30 * time.Hour),
			UpdatedAt:   time.Now().Add(-5 * time.Hour),
		},
		{
			ID:          "12",
			Title:       "Security audit",
			Description: "Run penetration tests and fix vulnerabilities",
			Status:      "done",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-120 * time.Hour),
			UpdatedAt:   time.Now().Add(-48 * time.Hour),
		},
	}

	for _, task := range demoTasks {
		s.tasks[task.ID] = task
	}
	s.nextID = 13
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
