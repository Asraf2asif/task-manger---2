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
const sortTasks = (taskList, field, direction) => {
  const sorted = [...taskList];

  switch (field) {
    case "priority":
      return sorted.sort((a, b) => {
        const priorityOrder = { high: 0, medium: 1, low: 2 };
        const comparison =
          priorityOrder[a.priority] - priorityOrder[b.priority];
        return direction === "desc" ? comparison : -comparison;
      });
    case "name":
      return sorted.sort((a, b) => {
        const comparison = a.title.localeCompare(b.title);
        return direction === "asc" ? comparison : -comparison;
      });
    case "date":
      return sorted.sort((a, b) => {
        const comparison = new Date(a.updatedAt) - new Date(b.updatedAt);
        return direction === "asc" ? comparison : -comparison;
      });
    default:
      return sorted;
  }
};

function TaskList({ tasks, loading, error, onEdit, onDelete }) {
  const [todoPage, setTodoPage] = useState(1);
  const [inProgressPage, setInProgressPage] = useState(1);
  const [donePage, setDonePage] = useState(1);

  // Separate sort state for each tab - field and direction (default: priority high to low)
  const [todoSortField, setTodoSortField] = useState("priority");
  const [todoSortDir, setTodoSortDir] = useState("desc");

  const [inProgressSortField, setInProgressSortField] = useState("priority");
  const [inProgressSortDir, setInProgressSortDir] = useState("desc");

  const [doneSortField, setDoneSortField] = useState("priority");
  const [doneSortDir, setDoneSortDir] = useState("desc");

  // Group and sort tasks by status with independent sorting - BEFORE early returns
  const todoTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "todo"),
        todoSortField,
        todoSortDir
      ),
    [tasks, todoSortField, todoSortDir]
  );
  const inProgressTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "in-progress"),
        inProgressSortField,
        inProgressSortDir
      ),
    [tasks, inProgressSortField, inProgressSortDir]
  );
  const doneTasks = useMemo(
    () =>
      sortTasks(
        tasks.filter((t) => t.status === "done"),
        doneSortField,
        doneSortDir
      ),
    [tasks, doneSortField, doneSortDir]
  );

  // Reset pages to 1 when tasks change (e.g., after search)
  React.useEffect(() => {
    setTodoPage(1);
    setInProgressPage(1);
    setDonePage(1);
  }, [tasks.length]);

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

  const renderSortControl = (field, setField, direction, setDirection) => {
    // Get label for direction button based on field
    const getDirectionLabel = () => {
      if (field === "priority") {
        return direction === "desc" ? "High to Low" : "Low to High";
      } else if (field === "date") {
        return direction === "desc" ? "Newest First" : "Oldest First";
      } else {
        return direction === "asc" ? "A to Z" : "Z to A";
      }
    };

    // Get icon rotation based on direction
    const getIconRotation = () => {
      if (field === "priority") {
        return direction === "desc" ? "rotate-180" : "";
      } else if (field === "date") {
        return direction === "desc" ? "rotate-180" : "";
      } else {
        return direction === "asc" ? "" : "rotate-180";
      }
    };

    return (
      <div className="inline-flex items-center gap-2 mb-6 bg-white p-3 rounded-lg shadow-sm border">
        <ArrowUpDown
          size={16}
          className="text-muted-foreground flex-shrink-0"
        />
        <span className="text-sm font-medium text-muted-foreground">Sort:</span>
        <Select value={field} onValueChange={setField}>
          <SelectTrigger className="w-[110px] h-9 border-0 bg-transparent hover:bg-accent focus:ring-0">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="priority">Priority</SelectItem>
            <SelectItem value="date">Date</SelectItem>
            <SelectItem value="name">Name</SelectItem>
          </SelectContent>
        </Select>
        <div className="h-6 w-px bg-border" />
        <Button
          variant="ghost"
          size="sm"
          onClick={() => setDirection(direction === "asc" ? "desc" : "asc")}
          className="gap-2 h-9 px-3 hover:bg-accent font-medium"
        >
          <ArrowUpDown
            size={14}
            className={`transition-transform ${getIconRotation()}`}
          />
          <span className="text-sm">{getDirectionLabel()}</span>
        </Button>
      </div>
    );
  };

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
              {renderSortControl(
                todoSortField,
                setTodoSortField,
                todoSortDir,
                setTodoSortDir
              )}
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
              {renderSortControl(
                inProgressSortField,
                setInProgressSortField,
                inProgressSortDir,
                setInProgressSortDir
              )}
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
              {renderSortControl(
                doneSortField,
                setDoneSortField,
                doneSortDir,
                setDoneSortDir
              )}
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
