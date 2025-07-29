package Usecases

import (
	"errors"
	"strings"
	"task_manager/Domain"
)

type TaskUsecase struct {
	Repo Domain.ITaskRepository
}

func NewTaskUsecase(repo Domain.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{Repo: repo}
}

func (uc *TaskUsecase) GetTaskByID(id string) (Domain.Task, error) {
	if strings.TrimSpace(id) == "" {
		return Domain.Task{}, errors.New("task ID cannot be empty")
	}
	return uc.Repo.GetTaskByID(id)
}


func (uc *TaskUsecase) GetAllTask() ([]Domain.Task, error) {
	return uc.Repo.GetAllTask()
}

func (uc *TaskUsecase) AddTask(task Domain.Task) error {
	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)

	if task.Title == "" {
		return errors.New("task title cannot be empty")
	}
	if task.Description == "" {
		return errors.New("task description cannot be empty")
	}

	return uc.Repo.AddTask(task)
}


func (uc *TaskUsecase) UpdatedTask(taskID string, updatedTask Domain.Task) error {
	if strings.TrimSpace(taskID) == "" {
		return errors.New("task ID cannot be empty")
	}

	updatedTask.Title = strings.TrimSpace(updatedTask.Title)
	updatedTask.Description = strings.TrimSpace(updatedTask.Description)

	if updatedTask.Title == "" {
		return errors.New("task title cannot be empty")
	}
	if updatedTask.Description == "" {
		return errors.New("task description cannot be empty")
	}

	return uc.Repo.UpdateTask(taskID, updatedTask)
}


func (uc *TaskUsecase) DeleteTask(taskID string) error {
	if strings.TrimSpace(taskID) == "" {
		return errors.New("task ID cannot be empty")
	}
	return uc.Repo.DeleteTask(taskID)
}
