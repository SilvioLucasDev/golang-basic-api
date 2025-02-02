package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", GetAllTasks)
		tasks.GET("/:id", GetTaskByID)
		tasks.POST("/", CreateTask)
		tasks.PATCH("/:id", UpdateTask)
		tasks.DELETE("/:id", DeleteTask)
	}
}
