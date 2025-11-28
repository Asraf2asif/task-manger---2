# Task Manager App

A clean task management application built with Go and React. Nothing fancy, just solid code that works.

## What's Inside

**Backend (Go):**

- REST API with proper error handling
- In-memory storage (no database needed for this demo)
- Clean separation of concerns across multiple files
- CORS enabled for local development

**Frontend (React + Material-UI):**

- Modern, responsive interface
- Task creation, editing, and deletion
- Tasks organized by status (To Do, In Progress, Done)
- Priority levels with color coding

## Getting Started

### Prerequisites

You'll need these installed:

- Go 1.21 or higher
- Node.js 16 or higher
- npm or yarn

### Running the Backend

Open a terminal and navigate to the backend folder:

```bash
cd backend
go mod download
go run cmd/server/main.go
```

The server starts on port 8080. You should see "Server starting on port 8080" in your terminal.

### Running the Frontend

**Using Vite with Shadcn/UI:**

```bash
cd frontend-vite
npm install

# Install shadcn components (run these one by one)
npx shadcn@latest add button
npx shadcn@latest add card
npx shadcn@latest add dialog
npx shadcn@latest add input
npx shadcn@latest add label
npx shadcn@latest add select
npx shadcn@latest add badge

# Start the dev server
npm run dev
```

Opens your browser at http://localhost:3000.

## How to Use It

1. Click "New Task" in the top right to create a task
2. Fill in the title (required), description, status, and priority
3. Click "Create" and your task appears in the list
4. Use the edit icon to modify a task
5. Use the delete icon to remove a task (it'll ask for confirmation)

Tasks are grouped by status so you can see what's in progress at a glance.

## Project Structure

```
backend/
  ├── cmd/
  │   └── server/
  │       └── main.go           - Entry point
  ├── internal/
  │   ├── models/
  │   │   └── task.go           - Data structures
  │   ├── storage/
  │   │   └── memory.go         - In-memory storage with demo data
  │   ├── handlers/
  │   │   ├── tasks.go          - Task request handlers
  │   │   └── utils.go          - Response helpers
  │   └── router/
  │       └── router.go         - Route configuration
  └── go.mod

frontend-vite/
  ├── src/
  │   ├── components/
  │   │   └── ui/               - Shadcn/UI components (auto-generated)
  │   ├── hooks/
  │   │   └── useTasks.js       - React Query hooks
  │   ├── lib/
  │   │   └── utils.js          - Utility functions
  │   ├── services/
  │   │   └── api.js            - API communication
  │   ├── App.jsx               - Main app component
  │   ├── main.jsx              - React entry point
  │   └── index.css             - Tailwind styles
  ├── components.json           - Shadcn config
  └── tailwind.config.js        - Tailwind config
```

## Notes

The backend uses demo data that's hardcoded in storage.go. When you restart the server, any changes you made will be lost. That's intentional - this is a demo, not a production app. In a real application, you'd connect to a database like PostgreSQL or MongoDB.

The frontend assumes the backend runs on localhost:8080. If you change the port, update the API_BASE constant in frontend/src/services/api.js.

## Troubleshooting

**Frontend can't connect to backend:**

- Make sure the backend is running on port 8080
- Check that CORS is working (it should be enabled by default)
- Look at the browser console for specific error messages

**Port already in use:**

- Change the port by setting the PORT environment variable before running the backend
- On Windows: `set PORT=3001 && go run .`
- On Mac/Linux: `PORT=3001 go run .`

That's it. The code is straightforward and each file does one thing. No magic, no over-engineering.
