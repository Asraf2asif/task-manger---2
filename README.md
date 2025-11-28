# 🚀 Modern Task Manager

A professional, full-stack task management application showcasing modern web development practices with Go and React.

## 📹 Demo

https://github.com/user-attachments/assets/e0c8c8e5-8f5e-4f3e-9f5e-8f5e8f5e8f5e

<video src="https://raw.githubusercontent.com/Asraf2asif/task-manger---2/refs/heads/main/Demonstration%20video.mp4" controls></video>

> **Note:** If the video doesn't play above, [click here to download and watch](https://raw.githubusercontent.com/Asraf2asif/task-manger---2/refs/heads/main/Demonstration%20video.mp4)

## ✨ Features Overview

### 🎨 Frontend (React + Shadcn/UI + Tailwind CSS)

**User Interface:**

- 🎯 **Clean, Modern Design** - Professional UI built with Shadcn/UI components and Tailwind CSS
- 📱 **Fully Responsive** - Seamless experience across desktop, tablet, and mobile devices
- 🌊 **Smooth Animations** - Framer Motion animations for delightful user interactions
- 🎨 **Thoughtful UX** - Every interaction is carefully designed for optimal user experience

**Task Management:**

- ✅ **Tabbed Organization** - Separate tabs for To Do, In Progress, and Done tasks
- 🔍 **Real-time Search** - Instant search with clear button for quick filtering
- 📊 **Smart Sorting** - Independent sorting per tab with 3 options:
  - Priority (High to Low / Low to High)
  - Date (Newest First / Oldest First)
  - Name (A to Z / Z to A)
- 📄 **Pagination** - Clean pagination with 10 tasks per page
- 🎯 **Priority Badges** - Visual priority indicators (High/Medium/Low) with color coding
- 📝 **Rich Task Cards** - Display title, description, priority, and last updated date

**User Actions:**

- ➕ **Create Tasks** - Beautiful modal dialog with form validation
- ✏️ **Edit Tasks** - Inline editing with pre-populated forms
- 🗑️ **Delete Tasks** - Confirmation dialog to prevent accidental deletions
- 🔔 **Toast Notifications** - Real-time feedback for all actions (success/error)
- ⚡ **Instant Updates** - React Query for optimistic updates and caching

**Technical Highlights:**

- 🔄 **React Query** - Advanced state management with automatic caching and refetching
- 🎭 **Framer Motion** - Smooth page transitions and micro-interactions
- 🎨 **Shadcn/UI** - Accessible, customizable component library
- 🎯 **Lucide Icons** - Modern, consistent iconography
- 📦 **Vite** - Lightning-fast build tool and dev server

### ⚙️ Backend (Go + SQLite)

**API Architecture:**

- 🏗️ **Clean Architecture** - Well-organized folder structure with separation of concerns
- 🔌 **RESTful API** - Standard HTTP methods (GET, POST, PUT, DELETE)
- 🛡️ **Error Handling** - Comprehensive error handling with meaningful messages
- 📝 **Logging** - Structured logging for debugging and monitoring
- 🌐 **CORS Enabled** - Configured for local development and production

**Database:**

- 💾 **SQLite Storage** - Persistent storage with pure Go implementation (no CGO required)
- 🔄 **Auto-Migration** - Database schema created automatically on first run
- 🌱 **Demo Data** - 72 realistic demo tasks seeded on initialization
- 🔍 **Indexed Queries** - Optimized with database indexes for fast lookups
- 🔒 **Thread-Safe** - Concurrent request handling with proper locking

**Code Quality:**

- 📁 **Modular Design** - Clean separation: models, storage, handlers, router
- 🔌 **Interface-Based** - Storage interface allows easy swapping of implementations
- 📏 **Small Files** - No file exceeds 200 lines for maintainability
- 💬 **Well Commented** - Clear, concise comments explaining business logic
- 🧪 **Production Ready** - Proper error handling, validation, and edge case coverage

## 🛠️ Tech Stack

**Frontend:**

- React 18 - Modern UI library
- Vite - Next-generation build tool
- Shadcn/UI - Accessible component library
- Tailwind CSS - Utility-first CSS framework
- Framer Motion - Animation library
- React Query - Server state management
- React Hot Toast - Toast notifications
- Lucide React - Icon library

**Backend:**

- Go 1.21+ - Fast, compiled language
- Gorilla Mux - HTTP router
- modernc.org/sqlite - Pure Go SQLite driver
- Standard library - Minimal dependencies

## 📦 Project Structure

```
task-manager/
├── backend/
│   ├── internal/
│   │   ├── models/          # Data structures
│   │   ├── storage/         # Database layer
│   │   │   ├── store.go     # Storage interface
│   │   │   └── sqlite.go    # SQLite implementation
│   │   ├── handlers/        # HTTP request handlers
│   │   │   ├── tasks.go     # Task CRUD operations
│   │   │   └── utils.go     # Response helpers
│   │   └── router/          # Route configuration
│   ├── main.go              # Application entry point
│   ├── tasks.db             # SQLite database (auto-created)
│   └── go.mod               # Go dependencies
│
└── frontend-vite/
    ├── src/
    │   ├── components/
    │   │   ├── ui/          # Shadcn components
    │   │   ├── Header.jsx   # Top navigation with search
    │   │   ├── TaskList.jsx # Tabbed task display with sorting
    │   │   ├── TaskCard.jsx # Individual task card
    │   │   ├── TaskForm.jsx # Create/edit modal
    │   │   └── DeleteDialog.jsx # Confirmation dialog
    │   ├── hooks/
    │   │   └── useTasks.js  # React Query hooks
    │   ├── services/
    │   │   └── api.js       # API client
    │   ├── lib/
    │   │   └── utils.js     # Utility functions
    │   ├── App.jsx          # Main component
    │   └── main.jsx         # React entry point
    ├── components.json      # Shadcn configuration
    └── tailwind.config.js   # Tailwind configuration
```

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** - [Download](https://golang.org/dl/)
- **Node.js 16+** - [Download](https://nodejs.org/)
- **npm or yarn** - Comes with Node.js

### Installation & Running

**1. Clone the repository**

```bash
git clone <your-repo-url>
cd task-manager
```

**2. Start the Backend**

```bash
cd backend
go run .
```

The server starts on `http://localhost:8080`. On first run, it creates `tasks.db` and seeds 72 demo tasks.

**3. Start the Frontend** (in a new terminal)

```bash
cd frontend-vite
npm install
npm run dev
```

The app opens at `http://localhost:3000`

## 🎯 API Endpoints

| Method | Endpoint         | Description       |
| ------ | ---------------- | ----------------- |
| GET    | `/api/tasks`     | Get all tasks     |
| POST   | `/api/tasks`     | Create a new task |
| PUT    | `/api/tasks/:id` | Update a task     |
| DELETE | `/api/tasks/:id` | Delete a task     |
| GET    | `/api/health`    | Health check      |

## 💡 Key Features Explained

### Smart Sorting System

Each tab (To Do, In Progress, Done) has independent sorting:

- Select field: Priority, Date, or Name
- Toggle direction with visual feedback
- Sorting persists per tab

### Search with Auto-Reset

- Real-time filtering across all tasks
- Clear button appears when searching
- Automatically resets pagination to page 1

### Optimistic Updates

React Query provides instant UI updates while syncing with the server in the background.

### Confirmation Dialogs

Delete actions show a beautiful confirmation dialog instead of browser alerts.

### Responsive Pagination

- 10 tasks per page for optimal viewing
- Independent pagination per tab
- Smooth transitions between pages

## 🎨 Design Philosophy

This project demonstrates:

- **User-First Design** - Every feature serves a clear user need
- **Performance** - Fast load times, smooth animations, optimized queries
- **Accessibility** - Semantic HTML, ARIA labels, keyboard navigation
- **Maintainability** - Clean code, small files, clear structure
- **Scalability** - Interface-based design, modular architecture

## 📝 Notes

- Database file (`tasks.db`) persists data across restarts
- Demo tasks are seeded only on first run
- All changes are saved immediately to the database
- Pure Go SQLite driver works on all platforms without CGO

## 🤝 Contributing

This is a portfolio project, but suggestions and feedback are welcome!

## 📄 License

MIT License - feel free to use this project for learning or as a template.

---

**Built with ❤️ to showcase modern full-stack development practices**
