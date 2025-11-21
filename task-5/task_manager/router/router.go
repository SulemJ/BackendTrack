package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func StarterH() {
	var router = gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.AddNewTask)
	router.Run("localhost:8080")
}
