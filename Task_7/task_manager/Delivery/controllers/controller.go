package controllers

import (
	"net/http"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskControllers struct {
	TaskUsecase Usecases.TaskUsecase
}
type UserControllers struct {
    UserUsecase *Usecases.UserUsecase 
}

func (tc *TaskControllers) HandleGetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
func (uc *UserControllers) HandleRegister(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := uc.UserUsecase.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"Message": "user register successfully"})
}

func (uc *UserControllers) HandleLogin(c *gin.Context) {
	var loginRequest Domain.LoginRequest
	if err:= c.ShouldBindJSON(&loginRequest); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalide Request format"})
	}
	token, err := uc.UserUsecase.Login(c.Request.Context(),loginRequest)
	if err != nil{
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"message":token})
	

}

func (tc *TaskControllers) HandleCreateTask(c *gin.Context){
	var task Domain.Task 
	if err:= c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid Request"})
		return 
	}
	err := tc.TaskUsecase.AddTask()


}