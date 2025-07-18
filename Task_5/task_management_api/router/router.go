package router

import (
	"github.com/gin-gonic/gin"
	"task_management_api/controllers"
	)

func Taskrouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", controllers.HandlerGetTasks)
	router.GET("/tasks/:id",controllers.HandlerGetTaskByID)
	router.POST("/tasks", controllers.HandleAddTask)
	router.PUT("/tasks/:id", controllers.HandlerUpdateTask)
	router.DELETE("/tasks/:id", controllers.HandlerDeleteTask)

	return router

}
