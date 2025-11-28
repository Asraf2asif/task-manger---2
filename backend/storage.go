package main

import (
	"sync"
	"time"
)

// In-memory storage - this is just demo data, not production-ready
// In a real app, you'd use a proper database
var (
	tasks     = make(map[string]*Task)
	taskMutex = &sync.RWMutex{}
	nextID    = 1
)

func init() {
	// Seed with some demo data so the app doesn't look empty
	// These are fake tasks just for demonstration
	seedDemoTasks()
}

// seedDemoTasks creates initial demo data
func seedDemoTasks() {
	demoTasks := []*Task{
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
			Status:      "todo",
			Priority:    "low",
			CreatedAt:   time.Now().Add(-36 * time.Hour),
			UpdatedAt:   time.Now().Add(-36 * time.Hour),
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
			Status:      "todo",
			Priority:    "high",
			CreatedAt:   time.Now().Add(-8 * time.Hour),
			UpdatedAt:   time.Now().Add(-8 * time.Hour),
		},
		{
			ID:          "9",
			Title:       "Design new landing page",
			Description: "Create mockups for the updated homepage",
			Status:      "in-progress",
			Priority:    "low",
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
	}
	
	for _, task := range demoTasks {
		tasks[task.ID] = task
	}
	nextID = 11
}

// getAllTasks returns all tasks (thread-safe)
func getAllTasks() []*Task {
	taskMutex.RLock()
	defer taskMutex.RUnlock()
	
	result := make([]*Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task)
	}
	return result
}

// getTaskByID retrieves a single task (thread-safe)
func getTaskByID(id string) (*Task, bool) {
	taskMutex.RLock()
	defer taskMutex.RUnlock()
	
	task, exists := tasks[id]
	return task, exists
}

// saveTask stores or updates a task (thread-safe)
func saveTask(task *Task) {
	taskMutex.Lock()
	defer taskMutex.Unlock()
	
	tasks[task.ID] = task
}

// removeTask deletes a task (thread-safe)
func removeTask(id string) bool {
	taskMutex.Lock()
	defer taskMutex.Unlock()
	
	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		return true
	}
	return false
}
