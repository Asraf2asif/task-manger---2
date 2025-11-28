package storage

import (
	"database/sql"
	"log"
	"taskapp/internal/models"
	"time"

	_ "modernc.org/sqlite"
)

// SQLiteStore handles SQLite database storage
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore creates a new SQLite storage instance
func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{db: db}
	
	// Initialize database schema
	if err := store.initSchema(); err != nil {
		return nil, err
	}

	// Seed demo data if database is empty
	if err := store.seedDemoTasks(); err != nil {
		return nil, err
	}

	return store, nil
}

// initSchema creates the tasks table
func (s *SQLiteStore) initSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL,
		priority TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_status ON tasks(status);
	CREATE INDEX IF NOT EXISTS idx_priority ON tasks(priority);
	`

	_, err := s.db.Exec(query)
	return err
}

// seedDemoTasks inserts 72 demo tasks if database is empty
func (s *SQLiteStore) seedDemoTasks() error {
	// Check if database already has tasks
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Printf("Database already has %d tasks, skipping seed", count)
		return nil
	}

	log.Println("Seeding database with 72 demo tasks...")

	now := time.Now()
	demoTasks := []struct {
		id          string
		title       string
		description string
		status      string
		priority    string
		hoursAgo    int
	}{
		// To Do tasks (24)
		{"1", "Review pull requests", "Check the pending PRs from the team", "todo", "high", 48},
		{"2", "Write unit tests", "Increase test coverage to at least 80%", "todo", "medium", 12},
		{"3", "Refactor authentication module", "Clean up the auth code and improve security", "todo", "high", 8},
		{"4", "Mobile app research", "Evaluate React Native vs Flutter for mobile version", "todo", "low", 4},
		{"5", "Update dependencies", "Upgrade all npm packages to latest versions", "todo", "medium", 24},
		{"6", "Create user onboarding flow", "Design and implement welcome screens for new users", "todo", "high", 16},
		{"7", "Add email notifications", "Send email alerts for important task updates", "todo", "low", 6},
		{"8", "Implement search filters", "Add advanced filtering options for task search", "todo", "medium", 10},
		{"9", "Add export to CSV feature", "Allow users to export task lists to CSV", "todo", "low", 14},
		{"10", "Create mobile responsive design", "Optimize UI for mobile devices", "todo", "high", 18},
		{"11", "Implement task templates", "Create reusable task templates", "todo", "medium", 22},
		{"12", "Add task comments", "Enable commenting on tasks", "todo", "medium", 26},
		{"13", "Setup backup system", "Automated daily backups", "todo", "high", 30},
		{"14", "Add task dependencies", "Link tasks that depend on each other", "todo", "low", 34},
		{"15", "Create admin dashboard", "Build admin panel for user management", "todo", "high", 38},
		{"16", "Implement task reminders", "Send reminders for upcoming deadlines", "todo", "medium", 42},
		{"17", "Add keyboard shortcuts", "Implement hotkeys for common actions", "todo", "low", 46},
		{"18", "Create API webhooks", "Allow external integrations via webhooks", "todo", "medium", 50},
		{"19", "Add task labels/tags", "Categorize tasks with custom labels", "todo", "medium", 54},
		{"20", "Implement bulk operations", "Edit multiple tasks at once", "todo", "low", 58},
		{"21", "Add task history", "Track all changes made to tasks", "todo", "medium", 62},
		{"22", "Create calendar view", "Display tasks in calendar format", "todo", "high", 66},
		{"23", "Add time tracking", "Track time spent on each task", "todo", "medium", 70},
		{"24", "Implement task priorities", "Auto-sort by priority levels", "todo", "low", 74},

		// In Progress tasks (24)
		{"25", "Update API documentation", "Document the new endpoints we added last week", "in-progress", "medium", 24},
		{"26", "Implement dark mode", "Add theme toggle for better user experience", "in-progress", "low", 36},
		{"27", "Optimize database queries", "Some queries are taking too long, need indexing", "in-progress", "high", 20},
		{"28", "Design new landing page", "Create mockups for the updated homepage", "in-progress", "medium", 48},
		{"29", "Performance monitoring setup", "Integrate New Relic or DataDog for monitoring", "in-progress", "medium", 30},
		{"30", "Build analytics dashboard", "Create charts and metrics for user activity", "in-progress", "high", 40},
		{"31", "Implement file upload feature", "Allow users to attach files to tasks", "in-progress", "medium", 28},
		{"32", "Add real-time collaboration", "Enable multiple users to work on tasks simultaneously", "in-progress", "high", 32},
		{"33", "Build notification system", "Push notifications for task updates", "in-progress", "high", 44},
		{"34", "Create user profiles", "Add customizable user profile pages", "in-progress", "medium", 52},
		{"35", "Implement OAuth providers", "Add Google and GitHub login", "in-progress", "high", 38},
		{"36", "Add task attachments", "Allow file attachments to tasks", "in-progress", "medium", 46},
		{"37", "Create team workspaces", "Separate workspaces for different teams", "in-progress", "high", 50},
		{"38", "Implement task sorting", "Multiple sort options for task lists", "in-progress", "low", 42},
		{"39", "Add recurring tasks", "Tasks that repeat on schedule", "in-progress", "medium", 54},
		{"40", "Build reporting system", "Generate task completion reports", "in-progress", "medium", 58},
		{"41", "Create task board view", "Kanban-style board for tasks", "in-progress", "high", 62},
		{"42", "Add task milestones", "Group tasks into project milestones", "in-progress", "medium", 66},
		{"43", "Implement task assignments", "Assign tasks to specific users", "in-progress", "high", 70},
		{"44", "Add task subtasks", "Break down tasks into smaller subtasks", "in-progress", "medium", 74},
		{"45", "Create activity feed", "Show recent activity across all tasks", "in-progress", "low", 78},
		{"46", "Build custom fields", "Add custom metadata fields to tasks", "in-progress", "medium", 82},
		{"47", "Implement task templates", "Predefined task structures", "in-progress", "low", 86},
		{"48", "Add task automation", "Automate repetitive task actions", "in-progress", "high", 90},

		// Done tasks (24)
		{"49", "Fix login bug", "Users reported issues with OAuth login flow", "done", "high", 72},
		{"50", "Setup CI/CD pipeline", "Automate deployment process with GitHub Actions", "done", "medium", 96},
		{"51", "Security audit", "Run penetration tests and fix vulnerabilities", "done", "high", 120},
		{"52", "Migrate to PostgreSQL", "Move from SQLite to PostgreSQL for production", "done", "high", 144},
		{"53", "Add password reset flow", "Implement forgot password functionality", "done", "medium", 80},
		{"54", "Create API rate limiting", "Prevent API abuse with rate limits", "done", "high", 100},
		{"55", "Setup error tracking", "Integrate Sentry for error monitoring", "done", "medium", 88},
		{"56", "Implement caching layer", "Add Redis for caching frequently accessed data", "done", "high", 110},
		{"57", "Setup load balancer", "Configure nginx for load balancing", "done", "high", 130},
		{"58", "Add SSL certificates", "Setup HTTPS with Let's Encrypt", "done", "high", 140},
		{"59", "Create backup strategy", "Implement automated backup system", "done", "high", 150},
		{"60", "Setup monitoring alerts", "Configure PagerDuty alerts", "done", "medium", 160},
		{"61", "Implement logging system", "Centralized logging with ELK stack", "done", "high", 170},
		{"62", "Add input validation", "Validate all user inputs", "done", "high", 180},
		{"63", "Create API documentation", "Complete API docs with Swagger", "done", "medium", 190},
		{"64", "Setup staging environment", "Create staging server for testing", "done", "high", 200},
		{"65", "Implement CORS policy", "Configure proper CORS headers", "done", "medium", 210},
		{"66", "Add request throttling", "Prevent API abuse with throttling", "done", "high", 220},
		{"67", "Create health check endpoint", "Monitor service health", "done", "medium", 230},
		{"68", "Setup database indexes", "Optimize database performance", "done", "high", 240},
		{"69", "Implement session management", "Secure session handling", "done", "high", 250},
		{"70", "Add CSRF protection", "Prevent cross-site request forgery", "done", "high", 260},
		{"71", "Create deployment scripts", "Automate deployment process", "done", "medium", 270},
		{"72", "Setup container orchestration", "Deploy with Kubernetes", "done", "high", 280},
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO tasks (id, title, description, status, priority, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, task := range demoTasks {
		createdAt := now.Add(-time.Duration(task.hoursAgo) * time.Hour)
		updatedAt := createdAt.Add(time.Duration(task.hoursAgo/2) * time.Hour)

		_, err = stmt.Exec(
			task.id,
			task.title,
			task.description,
			task.status,
			task.priority,
			createdAt,
			updatedAt,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Printf("Successfully seeded 72 demo tasks")
	return nil
}

// GetAll returns all tasks
func (s *SQLiteStore) GetAll() []*models.Task {
	rows, err := s.db.Query(`
		SELECT id, title, description, status, priority, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return []*models.Task{}
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning task: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// GetByID retrieves a single task
func (s *SQLiteStore) GetByID(id string) (*models.Task, bool) {
	task := &models.Task{}
	err := s.db.QueryRow(`
		SELECT id, title, description, status, priority, created_at, updated_at
		FROM tasks WHERE id = ?
	`, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, false
	}
	if err != nil {
		log.Printf("Error fetching task: %v", err)
		return nil, false
	}

	return task, true
}

// Save stores or updates a task
func (s *SQLiteStore) Save(task *models.Task) {
	_, err := s.db.Exec(`
		INSERT OR REPLACE INTO tasks (id, title, description, status, priority, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.CreatedAt,
		task.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error saving task: %v", err)
	}
}

// Delete removes a task
func (s *SQLiteStore) Delete(id string) bool {
	result, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return false
	}

	rows, _ := result.RowsAffected()
	return rows > 0
}

// GetNextID returns the next available ID
func (s *SQLiteStore) GetNextID() int {
	var maxID int
	err := s.db.QueryRow("SELECT COALESCE(MAX(CAST(id AS INTEGER)), 0) FROM tasks").Scan(&maxID)
	if err != nil {
		log.Printf("Error getting next ID: %v", err)
		return 1
	}
	return maxID + 1
}

// Close closes the database connection
func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
