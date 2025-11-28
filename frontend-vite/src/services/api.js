// API base URL - change this if your backend runs on a different port
const API_BASE = "http://localhost:8080/api";

// Helper function to handle API responses
const handleResponse = async (response) => {
  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ message: "Request failed" }));
    throw new Error(error.message || "Something went wrong");
  }

  // Handle 204 No Content responses
  if (response.status === 204) {
    return null;
  }

  return response.json();
};

// Get all tasks
export const getTasks = async () => {
  const response = await fetch(`${API_BASE}/tasks`);
  return handleResponse(response);
};

// Create a new task
export const createTask = async (taskData) => {
  const response = await fetch(`${API_BASE}/tasks`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(taskData),
  });
  return handleResponse(response);
};

// Update an existing task
export const updateTask = async (id, taskData) => {
  const response = await fetch(`${API_BASE}/tasks/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(taskData),
  });
  return handleResponse(response);
};

// Delete a task
export const deleteTask = async (id) => {
  const response = await fetch(`${API_BASE}/tasks/${id}`, {
    method: "DELETE",
  });
  return handleResponse(response);
};
