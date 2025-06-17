package main

import (
	"github/santori/tasks/handlers"
	"github/santori/tasks/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	taskStore := store.NewTaskStore()
	taskHandler := handlers.NewTaskHandler(taskStore)
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTask)

	r.Run(":8080")
}
