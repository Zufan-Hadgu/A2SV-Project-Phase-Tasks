package Usecases_test

import (
	"task_manager/Domain"
	"task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"task_manager/mocks"

)

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUsecase(mockRepo)

	expected := Domain.Task {Title: "Test Task", Description: "Test Desc"}
	mockRepo.On("GetTaskByID", "123").Return(expected, nil)

	result, err := usecase.GetTaskByID("123")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_EmptyID(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUsecase(mockRepo)

	_, err := usecase.GetTaskByID("   ")

	assert.NotNil(t, err)
	assert.Equal(t, "task ID cannot be empty", err.Error())
}

func TestAddTask_EmptyTitle(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUsecase(mockRepo)

	task := Domain.Task{Title: "", Description: "Valid"}
	err := usecase.AddTask(task)

	assert.NotNil(t, err)
	assert.Equal(t, "task title cannot be empty", err.Error())
}
