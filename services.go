package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetAllTasks(c *gin.Context) {
	rows, err := DB.Query("SELECT id, title FROM tasks")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch tasks",
		})
		return
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to fetch tasks",
			})
			return
		}
		tasks = append(tasks, task)
	}

	if tasks == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No tasks found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task Task

	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	if err := row.Scan(&task.ID, &task.Title); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func CreateTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create task",
		})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create task",
		})
		return
	}

	newTask.ID = int(id)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    newTask,
	})
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec("UPDATE tasks SET title = ? WHERE id = ?", updatedTask.Title, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
