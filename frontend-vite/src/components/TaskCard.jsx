import React, { useState } from "react";
import { Edit2, Trash2, AlertCircle, Clock } from "lucide-react";
import { motion } from "framer-motion";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import DeleteDialog from "./DeleteDialog.jsx";

function TaskCard({ task, onEdit, onDelete }) {
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);

  // Priority configuration
  const priorityConfig = {
    low: { variant: "secondary", icon: Clock },
    medium: { variant: "default", icon: AlertCircle },
    high: { variant: "destructive", icon: AlertCircle },
  };

  const config = priorityConfig[task.priority];
  const PriorityIcon = config.icon;

  const handleDeleteConfirm = () => {
    onDelete(task.id);
    setDeleteDialogOpen(false);
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, scale: 0.9 }}
      whileHover={{ y: -4 }}
      transition={{ duration: 0.2 }}
    >
      <Card className="h-full flex flex-col hover:shadow-lg transition-shadow">
        <CardContent className="flex-1 pt-6">
          <div className="flex items-start justify-between gap-2 mb-3">
            <h3 className="font-semibold text-lg flex-1">{task.title}</h3>
            <Badge variant={config.variant} className="flex items-center gap-1">
              <PriorityIcon size={12} />
              {task.priority}
            </Badge>
          </div>
          <p className="text-sm text-muted-foreground mb-4">
            {task.description || "No description"}
          </p>
          <p className="text-xs text-muted-foreground">
            Updated {new Date(task.updatedAt).toLocaleDateString()}
          </p>
        </CardContent>
        <CardFooter className="gap-2 pt-0">
          <motion.div whileHover={{ scale: 1.1 }} whileTap={{ scale: 0.9 }}>
            <Button
              variant="ghost"
              size="icon"
              onClick={() => onEdit(task)}
              aria-label="edit task"
            >
              <Edit2 size={18} />
            </Button>
          </motion.div>
          <motion.div whileHover={{ scale: 1.1 }} whileTap={{ scale: 0.9 }}>
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setDeleteDialogOpen(true)}
              aria-label="delete task"
              className="text-destructive hover:text-destructive"
            >
              <Trash2 size={18} />
            </Button>
          </motion.div>
        </CardFooter>
      </Card>
      <DeleteDialog
        open={deleteDialogOpen}
        onOpenChange={setDeleteDialogOpen}
        onConfirm={handleDeleteConfirm}
        taskTitle={task.title}
      />
    </motion.div>
  );
}

export default TaskCard;
