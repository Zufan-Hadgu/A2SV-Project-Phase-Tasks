package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_management_api/models"
	"github.com/zaahidali/task_management_api/data"
)


var TaskService data.TManager = *data.NewTManager()

func HandleAddTask(T *gin.Context){
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
	T.JSON(http.StatusCreated,gin.H{"message":" Task Added successfuly"})
}

func  HandlerGetTasks(T *gin.Context){

	Tasks := TaskService.GetAllTasks()
	T.JSON(http.StatusOK,Tasks)
}
func HandlerGetTaskByID(T *gin.Context){
	id := T.Param("id")
	task,err := TaskService.GetTaskByID(id)
	if err != nil {
		T.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	T.JSON(http.StatusOK,task)

}

func HandlerUpdateTask(T *gin.Context){
	id := T.Param("id")
	var updatedTask models.Task
	if err := T.ShouldBindJSON(&updatedTask); err != nil{
		T.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	err := TaskService.UpdateTask(id,updatedTask)
	if err != nil{
		T.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
		
	T.JSON(http.StatusOK, gin.H{"Message": "Task updated"})

	

}

func HandlerDeleteTask(T *gin.Context){
	id := T.Param("id")
	err := TaskService.DeteleTask(id)
	if err != nil{
		T.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	T.JSON(http.StatusOK, gin.H{"Message": "Task Deleted"})




}