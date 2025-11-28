import React, { useState, useMemo } from "react";
import { Toaster } from "react-hot-toast";
import Header from "./components/Header.jsx";
import TaskList from "./components/TaskList.jsx";
import TaskForm from "./components/TaskForm.jsx";
import {
  useTasks,
  useCreateTask,
  useUpdateTask,
  useDeleteTask,
} from "./hooks/useTasks.js";

function App() {
  const [formOpen, setFormOpen] = useState(false);
  const [editingTask, setEditingTask] = useState(null);
  const [searchQuery, setSearchQuery] = useState("");

  // React Query hooks
  const { data: tasks = [], isLoading, error } = useTasks();
  const createMutation = useCreateTask();
  const updateMutation = useUpdateTask();
  const deleteMutation = useDeleteTask();

  // Filter tasks based on search query
  const filteredTasks = useMemo(() => {
    if (!searchQuery.trim()) return tasks;

    const query = searchQuery.toLowerCase();
    return tasks.filter(
      (task) =>
        task.title.toLowerCase().includes(query) ||
        task.description.toLowerCase().includes(query)
    );
  }, [tasks, searchQuery]);

  const handleCreateTask = async (taskData) => {
    await createMutation.mutateAsync(taskData);
    setFormOpen(false);
  };

  const handleUpdateTask = async (taskData) => {
    await updateMutation.mutateAsync({ id: editingTask.id, data: taskData });
    setEditingTask(null);
    setFormOpen(false);
  };

  const handleDeleteTask = async (taskId) => {
    await deleteMutation.mutateAsync(taskId);
  };

  const handleEditClick = (task) => {
    setEditingTask(task);
    setFormOpen(true);
  };

  const handleFormClose = () => {
    setFormOpen(false);
    setEditingTask(null);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Header
        onAddClick={() => setFormOpen(true)}
        searchQuery={searchQuery}
        onSearchChange={setSearchQuery}
      />
      <div className="container mx-auto px-4 py-8 max-w-7xl">
        <TaskList
          tasks={filteredTasks}
          loading={isLoading}
          error={error?.message}
          onEdit={handleEditClick}
          onDelete={handleDeleteTask}
        />
      </div>
      <TaskForm
        open={formOpen}
        task={editingTask}
        onClose={handleFormClose}
        onSubmit={editingTask ? handleUpdateTask : handleCreateTask}
        isLoading={createMutation.isPending || updateMutation.isPending}
      />
      <Toaster position="top-right" />
    </div>
  );
}

export default App;
