package mocks

import (
	"github.com/stretchr/testify/mock"
	"task_manager/Domain"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTaskByID(id string) (Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAllTask() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) AddTask(task Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) UpdateTask(id string, task Domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// ------------------ User Repo ------------------

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByUsername(username string) (*Domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Domain.User), args.Error(1)
}


func (m *MockUserRepository) CountDB() (int64, error) {
    args := m.Called()
    return args.Get(0).(int64), args.Error(1)
}


func (m *MockUserRepository) Create(user Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// ------------------ Password Service ------------------

type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) ComparePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

// ------------------ JWT Service ------------------

type MockJwtService struct {
	mock.Mock
}

func (m *MockJwtService) GenerateToken(username, role string) (string, error) {
	args := m.Called(username, role)
	return args.String(0), args.Error(1)
}
