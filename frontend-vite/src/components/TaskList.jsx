import React, { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { CheckCircle2, Clock, PlayCircle, Loader2 } from "lucide-react";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import TaskCard from "./TaskCard.jsx";

const TASKS_PER_PAGE = 6;

function TaskList({ tasks, loading, error, onEdit, onDelete }) {
  const [currentPage, setCurrentPage] = useState(1);

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

  // Group tasks by status
  const groupedTasks = {
    todo: tasks.filter((t) => t.status === "todo"),
    "in-progress": tasks.filter((t) => t.status === "in-progress"),
    done: tasks.filter((t) => t.status === "done"),
  };

  const statusConfig = {
    todo: { label: "To Do", icon: Clock, color: "text-blue-600" },
    "in-progress": {
      label: "In Progress",
      icon: PlayCircle,
      color: "text-orange-600",
    },
    done: { label: "Done", icon: CheckCircle2, color: "text-green-600" },
  };

  // Pagination logic
  const totalPages = Math.ceil(tasks.length / TASKS_PER_PAGE);
  const startIndex = (currentPage - 1) * TASKS_PER_PAGE;
  const endIndex = startIndex + TASKS_PER_PAGE;
  const paginatedTasks = tasks.slice(startIndex, endIndex);

  // Group paginated tasks by status
  const paginatedGrouped = {
    todo: paginatedTasks.filter((t) => t.status === "todo"),
    "in-progress": paginatedTasks.filter((t) => t.status === "in-progress"),
    done: paginatedTasks.filter((t) => t.status === "done"),
  };

  return (
    <div>
      <AnimatePresence mode="wait">
        {Object.entries(paginatedGrouped).map(
          ([status, statusTasks]) =>
            statusTasks.length > 0 && (
              <motion.div
                key={status}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                className="mb-8"
              >
                <div className="flex items-center gap-2 mb-4">
                  {React.createElement(statusConfig[status].icon, {
                    size: 24,
                    className: statusConfig[status].color,
                  })}
                  <h2
                    className={`text-2xl font-bold ${statusConfig[status].color}`}
                  >
                    {statusConfig[status].label} ({groupedTasks[status].length})
                  </h2>
                </div>
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                  <AnimatePresence>
                    {statusTasks.map((task, index) => (
                      <motion.div
                        key={task.id}
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ delay: index * 0.05 }}
                      >
                        <TaskCard
                          task={task}
                          onEdit={onEdit}
                          onDelete={onDelete}
                        />
                      </motion.div>
                    ))}
                  </AnimatePresence>
                </div>
              </motion.div>
            )
        )}
      </AnimatePresence>

      {totalPages > 1 && (
        <div className="flex justify-center items-center gap-2 mt-8">
          <Button
            variant="outline"
            onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
            disabled={currentPage === 1}
          >
            Previous
          </Button>
          <span className="text-sm text-muted-foreground">
            Page {currentPage} of {totalPages}
          </span>
          <Button
            variant="outline"
            onClick={() => setCurrentPage((p) => Math.min(totalPages, p + 1))}
            disabled={currentPage === totalPages}
          >
            Next
          </Button>
        </div>
      )}
    </div>
  );
}

export default TaskList;
