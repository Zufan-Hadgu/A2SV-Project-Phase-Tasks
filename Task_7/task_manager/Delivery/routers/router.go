package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	client := Infrastructure.ConnetDB("mongodb://localhost:27017/")
	// TaskCollection := client.Database("taskdb").Collection("tasks")
	UserCollection := client.Database("taskdb").Collection("users")

	// taskRepo := Repositories.NewTaskRepoImpl(TaskCollection)
	userRepo := Repositories.NewUserRepoImpl(UserCollection)

	// taskUC := Usecases.NewTaskUsecase(taskRepo)
	userUC := Usecases.NewUserUsecase(userRepo)
	

	// taskController := &controllers.TaskControllers{TaskUsecase: taskUC}
	userController := &controllers.UserControllers{UserUsecase: userUC}

	// r.GET("/tasks/:id", taskController.HandleGetTaskByID)
	router.POST("/register", userController.HandleRegister)
	router.POST("/login",userController.HandleLogin)

	return router
}