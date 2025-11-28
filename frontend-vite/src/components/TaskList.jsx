import React, { useState, useMemo } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
  CheckCircle2,
  Clock,
  PlayCircle,
  Loader2,
  ArrowUpDown,
} from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import TaskCard from "./TaskCard.jsx";

const TASKS_PER_PAGE = 10;

// Sort function outside component
const sortTasks = (taskList, sortBy) => {
  const sorted = [...taskList];

  switch (sortBy) {
    case "priority-high":
      return sorted.sort((a, b) => {
        const priorityOrder = { high: 0, medium: 1, low: 2 };
        return priorityOrder[a.priority] - priorityOrder[b.priority];
      });
    case "priority-low":
      return sorted.sort((a, b) => {
        const priorityOrder = { high: 2, medium: 1, low: 0 };
        return priorityOrder[a.priority] - priorityOrder[b.priority];
      });
    case "name-asc":
      return sorted.sort((a, b) => a.title.localeCompare(b.title));
    case "name-desc":
      return sorted.sort((a, b) => b.title.localeCompare(a.title));
    case "date-asc":
      return sorted.sort(
        (a, b) => new Date(a.updatedAt) - new Date(b.updatedAt)
      );
    case "date-desc":
      return sorted.sort(
        (a, b) => new Date(b.updatedAt) - new Date(a.updatedAt)
      );
    default:
      return sorted;
  }
};

function TaskList({ tasks, loading, error, onEdit, onDelete }) {
  const [todoPage, setTodoPage] = useState(1);
  const [inProgressPage, setInProgressPage] = useState(1);
  const [donePage, setDonePage] = useState(1);

  // Separate sort state for each tab
  const [todoSort, setTodoSort] = useState("date-desc");
  const [inProgressSort, setInProgressSort] = useState("date-desc");
  const [doneSort, setDoneSort] = useState("date-desc");

  // Group and sort tasks by status with independent sorting - BEFORE early returns
  const todoTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "todo"),
        todoSort
      ),
    [tasks, todoSort]
  );
  const inProgressTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "in-progress"),
        inProgressSort
      ),
    [tasks, inProgressSort]
  );
  const doneTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "done"),
        doneSort
      ),
    [tasks, doneSort]
  );

  // Early returns AFTER all hooks
  if (loading) {
    return (
      <div className="flex justify-center items-center py-16">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    );
  }

  if (error) {
    return (
      <Alert variant="destructive" className="mb-4">
        <AlertDescription>{error}</AlertDescription>
      </Alert>
    );
  }

  if (tasks.length === 0) {
    return (
      <div className="text-center py-16">
        <h3 className="text-xl font-semibold text-muted-foreground">
          No tasks found. Create one to get started!
        </h3>
      </div>
    );
  }

  // Pagination helper
  const paginateTasks = (taskList, currentPage) => {
    const totalPages = Math.ceil(taskList.length / TASKS_PER_PAGE);
    const startIndex = (currentPage - 1) * TASKS_PER_PAGE;
    const endIndex = startIndex + TASKS_PER_PAGE;
    const paginatedTasks = taskList.slice(startIndex, endIndex);
    return { paginatedTasks, totalPages };
  };

  const { paginatedTasks: paginatedTodo, totalPages: todoTotalPages } =
    paginateTasks(todoTasks, todoPage);
  const {
    paginatedTasks: paginatedInProgress,
    totalPages: inProgressTotalPages,
  } = paginateTasks(inProgressTasks, inProgressPage);
  const { paginatedTasks: paginatedDone, totalPages: doneTotalPages } =
    paginateTasks(doneTasks, donePage);

  const renderTaskGrid = (taskList) => (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <AnimatePresence>
        {taskList.map((task, index) => (
          <motion.div
            key={task.id}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: index * 0.05 }}
          >
            <TaskCard task={task} onEdit={onEdit} onDelete={onDelete} />
          </motion.div>
        ))}
      </AnimatePresence>
    </div>
  );

  const renderPagination = (currentPage, totalPages, setPage) => {
    if (totalPages <= 1) return null;

    return (
      <div className="flex justify-center items-center gap-2 mt-6">
        <Button
          variant="outline"
          size="sm"
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          disabled={currentPage === 1}
        >
          Previous
        </Button>
        <span className="text-sm text-muted-foreground">
          Page {currentPage} of {totalPages}
        </span>
        <Button
          variant="outline"
          size="sm"
          onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
          disabled={currentPage === totalPages}
        >
          Next
        </Button>
      </div>
    );
  };

  const renderSortControl = (sortValue, setSortValue) => (
    <div className="flex items-center justify-between mb-6 bg-white p-4 rounded-lg shadow-sm">
      <div className="flex items-center gap-2">
        <ArrowUpDown size={18} className="text-muted-foreground" />
        <span className="text-sm font-medium">Sort by:</span>
      </div>
      <Select value={sortValue} onValueChange={setSortValue}>
        <SelectTrigger className="w-[200px]">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="date-desc">Date (Newest First)</SelectItem>
          <SelectItem value="date-asc">Date (Oldest First)</SelectItem>
          <SelectItem value="priority-high">Priority (High to Low)</SelectItem>
          <SelectItem value="priority-low">Priority (Low to High)</SelectItem>
          <SelectItem value="name-asc">Name (A to Z)</SelectItem>
          <SelectItem value="name-desc">Name (Z to A)</SelectItem>
        </SelectContent>
      </Select>
    </div>
  );

  return (
    <div>
      <Tabs defaultValue="todo" className="w-full">
        <TabsList className="grid w-full max-w-md mx-auto grid-cols-3 mb-8">
          <TabsTrigger value="todo" className="flex items-center gap-2">
            <Clock size={16} />
            To Do ({todoTasks.length})
          </TabsTrigger>
          <TabsTrigger value="in-progress" className="flex items-center gap-2">
            <PlayCircle size={16} />
            In Progress ({inProgressTasks.length})
          </TabsTrigger>
          <TabsTrigger value="done" className="flex items-center gap-2">
            <CheckCircle2 size={16} />
            Done ({doneTasks.length})
          </TabsTrigger>
        </TabsList>

        <TabsContent value="todo">
          {todoTasks.length === 0 ? (
            <div className="text-center py-16">
              <p className="text-muted-foreground">No tasks to do</p>
            </div>
          ) : (
            <>
              {renderSortControl(todoSort, setTodoSort)}
              {renderTaskGrid(paginatedTodo)}
              {renderPagination(todoPage, todoTotalPages, setTodoPage)}
            </>
          )}
        </TabsContent>

        <TabsContent value="in-progress">
          {inProgressTasks.length === 0 ? (
            <div className="text-center py-16">
              <p className="text-muted-foreground">No tasks in progress</p>
            </div>
          ) : (
            <>
              {renderSortControl(inProgressSort, setInProgressSort)}
              {renderTaskGrid(paginatedInProgress)}
              {renderPagination(
                inProgressPage,
                inProgressTotalPages,
                setInProgressPage
              )}
            </>
          )}
        </TabsContent>

        <TabsContent value="done">
          {doneTasks.length === 0 ? (
            <div className="text-center py-16">
              <p className="text-muted-foreground">No completed tasks</p>
            </div>
          ) : (
            <>
              {renderSortControl(doneSort, setDoneSort)}
              {renderTaskGrid(paginatedDone)}
              {renderPagination(donePage, doneTotalPages, setDonePage)}
            </>
          )}
        </TabsContent>
      </Tabs>
    </div>
  );
}

export default TaskList;
