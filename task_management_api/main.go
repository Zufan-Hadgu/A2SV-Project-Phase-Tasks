package main

import (
	"github.com/zaahidali/task_management_api/router"
)
func main(){
	r := router.Taskrouter()
	r.Run()	
}
