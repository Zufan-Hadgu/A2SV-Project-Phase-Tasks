package controllers

import (
	"net/http"
	"task_manager/controllers/dto"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var TaskService *data.TManager

var UserService *data.ImplUser



func HandleRegistration(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid request"})
		return
	}
	err := UserService.Register(user)
	if err != nil{
		c.JSON(http.StatusBadGateway,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"successfuly registerd"})
}
func HandleLogin(c *gin.Context) {
    var loginInfo dto.LoginRequest
    if err := c.ShouldBindJSON(&loginInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    token, err := UserService.Login(loginInfo.Username, loginInfo.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func HandleAddTask( c *gin.Context){
	var newTask models.Task
	role,_ := c.Get("role")

	if role != "admin"{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Your are not authorized user"})
		return
	}
	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Request"})
			return 

	}

	addedTask := TaskService.AddTask(newTask)
	if addedTask != nil{
		c.JSON(http.StatusBadGateway,gin.H{"error":"Internal error"})
		return 

	}
	c.JSON(http.StatusOK,gin.H{"message":"Succesfully added"})


}

func HandleDeleteTask(c *gin.Context){
	role , _ := c.Get("role")
	id := c.Param("id")
	if role != "admin"{
		c.JSON(http.StatusBadGateway,gin.H{"error":"unathorized user"})

	}
	err := TaskService.DeletDeteleTask(id)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"Task deleted successfuly "})

}

func HadleUpdateTask(c *gin.Context){
	id := c.Param("id")
	role ,_ := c.Get("role")
	var updatedTask models.Task
	if role != "admin"{
		c.JSON(http.StatusBadGateway,gin.H{"error":"unathorized user"})

	}

	if err := c.ShouldBindJSON(&updatedTask); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Request"})
			return 

	}
	err := TaskService.UpdateTask(id,updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Task updated successfuly "})

}
func HandleGetTaskByID(c *gin.Context){
	id := c.Param("id")
	task,err := TaskService.GetTaskByID(id)
	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,task)
}

func HandleGetAllTasks(c *gin.Context){
	Tasks,err := TaskService.GetAllTasks()
	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,Tasks)


}




