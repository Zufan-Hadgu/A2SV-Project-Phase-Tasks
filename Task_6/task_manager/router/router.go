package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
	"task_manager/middleware"
)

func Taskrouter() *gin.Engine{
	router := gin.Default()

	router.GET("/tasks",controllers.HandleGetAllTasks)
	router.POST("/register",controllers.HandleRegistration)
	router.POST("login",controllers.HandleLogin)
	router.GET("/tasks/:id",controllers.HandleGetTaskByID)
	router.POST("/tasks",middleware.AuthMiddleware(),controllers.HandleAddTask)
	router.PUT("/tasks/:id",middleware.AuthMiddleware(),controllers.HadleUpdateTask)
	router.DELETE("/tasks/:id",middleware.AuthMiddleware(),controllers.HandleDeleteTask)
	return router

}