package Usecases

import (
	"context"
	"task_manager/Domain"
)

type TaskUsecase struct {
	Repo Domain.TaskRepository
}

func NewTaskUsecase(repo Domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{Repo: repo}
}

func (uc *TaskUsecase) GetTaskByID(ctx context.Context, id string) (Domain.Task, error) {
	return uc.Repo.GetTaskByID(ctx, id)
}

func ( uc *TaskUsecase) GetAllTask(ctx context.Context)([]Domain.Task,error){
	return uc.Repo.GetAllTask(ctx)
}
func (uc *TaskUsecase) AddTask(ctx context.Context,task Domain.Task) error{
	return uc.Repo.AddTask(ctx,task)
}

func (uc *TaskUsecase) UpdatedTask(ctx context.Context,taskID string,UpdatedTask Domain.Task) error{
	return uc.Repo.UpdateTask(ctx,taskID,UpdatedTask)
}

func (uc *TaskUsecase) DeleteTask(ctx context.Context,taskID string) error{
	return uc.Repo.DeleteTask(ctx,taskID)
}