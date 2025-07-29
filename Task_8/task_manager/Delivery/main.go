package main

import "task_manager/Delivery/routers"

func main() {
	m := routers.SetupRouter()
	m.Run()	
	

}