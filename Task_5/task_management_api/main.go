package main

import (
	"task_management_api/controllers"
	"task_management_api/data"
	"task_management_api/router"
)

func main(){
	client := data. ConnectMongodbURI("mongodb://localhost:27017/")
	collection := client.Database("taskdb").Collection("tasks")

	taskService := &data.TManager{Collection: collection}

	controllers.TaskService = taskService   
	r := router.Taskrouter()
	r.Run()	
}