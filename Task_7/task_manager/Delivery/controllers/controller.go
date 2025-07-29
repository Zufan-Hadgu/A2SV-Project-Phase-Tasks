package controllers

import (
	"log"
	"net/http"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskControllers struct {
	TaskUsecase *Usecases.TaskUsecase
}
type UserControllers struct {
	UserUsecase *Usecases.UserUsecase
}

func (tc *TaskControllers) HandleGetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(id) 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (uc *UserControllers) HandleRegister(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {  
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (uc *UserControllers) HandleLogin(c *gin.Context) {
	var loginRequest Domain.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request format"})
		return
	}
	token, err := uc.UserUsecase.Login(loginRequest) 
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (tc *TaskControllers) HandleCreateTask(c *gin.Context) {
	var task Domain.Task
	role, _ := c.Get("role")
	log.Println(role)

	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := tc.TaskUsecase.AddTask(task) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (tc *TaskControllers) HandleUpdateTask(c *gin.Context) {
	var updatedTask Domain.Task
	role, _ := c.Get("role")
	taskID := c.Param("id")

	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := tc.TaskUsecase.UpdatedTask(taskID, updatedTask) 
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

func (tc *TaskControllers) HandleDeleteTask(c *gin.Context) {
	role, _ := c.Get("role")
	taskID := c.Param("id")

	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}
	err := tc.TaskUsecase.DeleteTask(taskID) 
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

func (tc *TaskControllers) HandleGetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTask() 
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
