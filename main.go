package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var taskList = []Task{
	{ID: 1, Title: "Task 1"},
	{ID: 2, Title: "Task 2"},
	{ID: 3, Title: "Task 3"},
}

func main() {
	fmt.Print("Minha primeira API")

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"tasks": taskList,
		})
	})

	router.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, task := range taskList {
			if fmt.Sprint(task.ID) == id {
				c.JSON(http.StatusOK, gin.H{
					"task": task,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
	})

	router.POST("/tasks", func(c *gin.Context) {
		var task Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task.ID = len(taskList) + 1
		taskList = append(taskList, task)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Task created successfully",
			"task":    task,
		})
	})

	router.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		for index, task := range taskList {
			if fmt.Sprint(task.ID) == id {
				taskList = append(taskList[:index], taskList[index+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message": "Task deleted successfully",
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
	})

	router.PATCH("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updatedTask Task

		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for index, task := range taskList {
			if fmt.Sprint(task.ID) == id {
				updatedTask.ID = task.ID
				taskList[index] = updatedTask
				c.JSON(http.StatusOK, gin.H{
					"message": "Task updated successfully",
					"task":    updatedTask,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
	})

	router.Run(":3000")
}
