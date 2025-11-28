import React from "react";
import { Plus, Search, X } from "lucide-react";
import { motion } from "framer-motion";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

function Header({ onAddClick, searchQuery, onSearchChange }) {
  return (
    <header className="bg-primary text-primary-foreground shadow-sm">
      <div className="container mx-auto px-4 py-4 max-w-7xl">
        <div className="flex items-center gap-4">
          <h1 className="text-2xl font-bold flex-shrink-0">Task Manager</h1>

          <div className="hidden sm:flex flex-1 max-w-md relative">
            <Search
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              size={18}
            />
            <Input
              type="text"
              placeholder="Search tasks..."
              value={searchQuery}
              onChange={(e) => onSearchChange(e.target.value)}
              className="pl-10 pr-10 bg-white/10 border-white/20 text-white placeholder:text-white/60"
            />
            {searchQuery && (
              <button
                onClick={() => onSearchChange("")}
                className="absolute right-3 top-1/2 transform -translate-y-1/2 text-white/60 hover:text-white transition-colors"
                aria-label="Clear search"
              >
                <X size={18} />
              </button>
            )}
          </div>

          <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
            <Button
              onClick={onAddClick}
              variant="secondary"
              className="hidden sm:flex items-center gap-2"
            >
              <Plus size={18} />
              New Task
            </Button>
            <Button
              onClick={onAddClick}
              variant="secondary"
              size="icon"
              className="sm:hidden"
            >
              <Plus size={20} />
            </Button>
          </motion.div>
        </div>
      </div>
    </header>
  );
}

export default Header;
