package main

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	client := data.ConnetDB("mongodb://localhost:27017/")
	TaskCollection := client.Database("taskdb").Collection("tasks")
	UserCollection := client.Database("taskdb").Collection("users")

	taskService := &data.TManager{Collection: TaskCollection}
	controllers.TaskService = taskService

	userService := &data.ImplUser{Collection:*UserCollection}
	controllers.UserService = userService
	r := router.Taskrouter()
	r.Run()	


}