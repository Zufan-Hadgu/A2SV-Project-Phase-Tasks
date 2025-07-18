package controllers

import (
	"net/http"
	"task_management_api/data"
	"task_management_api/models"

	"github.com/gin-gonic/gin"
)

var TaskService *data.TManager

func HandleAddTask( T *gin.Context){
	var newTask models.Task
	
	if err := T.ShouldBindJSON(&newTask); err != nil{
		T.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err := TaskService.AddTask(newTask)
	if err != nil{
		T.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}
	T.JSON(http.StatusOK,gin.H{"Message":"Task Added"})	
}

func HandlerGetTaskByID( T *gin.Context){
	taskID := T.Param("id")
	task,err := TaskService.GetTaskByID(taskID)
	if err != nil{
		T.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return	
	}
	T.JSON(http.StatusOK,task)
}

func HandlerUpdateTask(T *gin.Context){
	id := T.Param("id")
	var updatedTask models.Task
	if err := T.ShouldBindJSON(&updatedTask); err != nil{
		T.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := TaskService.UpdateTask(id,updatedTask)
	if err != nil{
		T.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}
		
	T.JSON(http.StatusOK, gin.H{"Message": "Task updated"})

	

}

func HandlerDeleteTask(T *gin.Context){
	id := T.Param("id")
	err := TaskService.DeteleTask(id)
	if err != nil{
		T.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return

	}
	T.JSON(http.StatusOK, gin.H{"Message": "Task Deleted"})

}
func  HandlerGetTasks(T *gin.Context){

	Tasks,err := TaskService.GetAllTasks()
	if err != nil{
		T.JSON(http.StatusNotFound,gin.H{"message":err.Error()})
	}
	T.JSON(http.StatusOK,Tasks)
}
