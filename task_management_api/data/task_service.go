package data

import (
	"fmt"
	"errors"	
	"time"	
	"github.com/zaahidali/task_management_api/models"
)



type TManager struct{
	tasks []models.Task
}

func NewTManager() *TManager {
    return &TManager{
		tasks: []models.Task{
			{ID: "1", Title: "First_tast", Description: "Study Go", DueDate: time.Now(), Status: "Pending"},
			{ID: "2", Title: "second_tast", Description: "Study MongoDB", DueDate: time.Now(), Status: "Pending"},
			{ID: "3", Title: "Third_tast", Description: "Study Go", DueDate: time.Now(), Status: "Pending"},
			{ID: "4", Title: "Fourth_tast", Description: "Study Go", DueDate: time.Now(), Status: "Done"},
        },
        
    }
}

type TaskManager interface{
	GetAllTasks() []models.Task
	GetTaskByID(taskID string ) (models.Task,error)
	UpdateTask(taskID string, updatedTask models.Task) error
	DeteleTask(taskID string) error
	AddTask(newTask models.Task) error

}

func (t * TManager) GetAllTasks() []models.Task{
	return t.tasks
}


func (t * TManager) GetTaskByID(taskID string ) (models.Task,error){
	
	for _, task := range t.tasks {
		if task.ID == taskID{
			return task ,nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (t * TManager) UpdateTask(taskID string, updatedTask models.Task) error {
	
	for i, task := range t.tasks {
		if task.ID == taskID {
			if updatedTask.Title != "" {
				t.tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				t.tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != ""{
				t.tasks[i].Status = updatedTask.Status
			}
			return nil
		}

	}
	return errors.New("task not found")

}

func (t * TManager) DeteleTask(taskID string) error {
	
	for i, task := range t.tasks {
		if task.ID == taskID{
			t.tasks = append(t.tasks[:i], t.tasks[:i+1]...)
			return nil
		}
	}
	return errors.New("task not found")
	

}

func (t *TManager) AddTask(newTask models.Task) error {
	newTask.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	t.tasks = append(t.tasks, newTask)
	return nil
}