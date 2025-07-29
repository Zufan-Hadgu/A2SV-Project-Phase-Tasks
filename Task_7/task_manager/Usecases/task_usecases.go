package Usecases

import (

	"task_manager/Domain"
	
)

type TaskUsecase struct {
	Repo Domain.ITaskRepository
}

func NewTaskUsecase(repo Domain.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{Repo: repo}
}

func (uc *TaskUsecase) GetTaskByID(id string) (Domain.Task, error) {
	return uc.Repo.GetTaskByID( id)
}

func ( uc *TaskUsecase) GetAllTask()([]Domain.Task,error){
	return uc.Repo.GetAllTask()
}
func (uc *TaskUsecase) AddTask(task Domain.Task) error{
	return uc.Repo.AddTask(task)
}

func (uc *TaskUsecase) UpdatedTask(taskID string,UpdatedTask Domain.Task) error{
	return uc.Repo.UpdateTask(taskID,UpdatedTask)
}

func (uc *TaskUsecase) DeleteTask(taskID string) error{
	return uc.Repo.DeleteTask(taskID)
}