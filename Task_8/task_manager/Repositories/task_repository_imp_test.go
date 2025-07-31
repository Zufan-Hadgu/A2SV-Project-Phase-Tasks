package Repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"task_manager/Domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ------------------------ Mock Definitions ------------------------

type MockCollection struct {
	mock.Mock
}

type CollectionInterface interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

// ------------------------ Test Suite ------------------------

type TaskRepoTestSuite struct {
	suite.Suite
	mockCol *MockCollection
	repo    *TaskRepoImpl
}

func (suite *TaskRepoTestSuite) SetupTest() {
	suite.mockCol = new(MockCollection)
	suite.repo = &TaskRepoImpl{Collection: suite.mockCol}
}

// ------------------------ Unit Tests ------------------------

func (suite *TaskRepoTestSuite) TestAddTask_Success() {
	task := Domain.Task{
		Title:       "Test Task",
		Description: "Test Desc",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	suite.mockCol.On("InsertOne", mock.Anything, task).Return(&mongo.InsertOneResult{}, nil)

	err := suite.repo.AddTask(task)

	assert.NoError(suite.T(), err)
	suite.mockCol.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestAddTask_Error() {
	task := Domain.Task{}
	suite.mockCol.On("InsertOne", mock.Anything, task).Return(nil, errors.New("insertion error"))

	err := suite.repo.AddTask(task)

	assert.EqualError(suite.T(), err, "error while creating task")
	suite.mockCol.AssertExpectations(suite.T())
}

// You would need a custom SingleResult mock to decode
// For brevity, we can focus on others like Update/Delete

func (suite *TaskRepoTestSuite) TestUpdateTask_Success() {
	taskID := "64cfe3f5a2165a7f621e81b4" // example valid ObjectID
	updatedTask := Domain.Task{
		Title:       "Updated",
		Description: "Updated Desc",
		DueDate:     time.Now(),
		Status:      "done",
	}
	filter, _ := BuildIDFilter(taskID)
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "description", Value: updatedTask.Description},
			{Key: "due_date", Value: updatedTask.DueDate},
			{Key: "status", Value: updatedTask.Status},
		},
	}}

	suite.mockCol.On("UpdateOne", mock.Anything, filter, update).Return(&mongo.UpdateResult{}, nil)

	err := suite.repo.UpdateTask(taskID, updatedTask)

	assert.NoError(suite.T(), err)
	suite.mockCol.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestDeleteTask_Success() {
	taskID := "64cfe3f5a2165a7f621e81b4"
	filter, _ := BuildIDFilter(taskID)

	suite.mockCol.On("DeleteOne", mock.Anything, filter).Return(&mongo.DeleteResult{}, nil)

	err := suite.repo.DeleteTask(taskID)

	assert.NoError(suite.T(), err)
	suite.mockCol.AssertExpectations(suite.T())
}

// ------------------------ Run the Suite ------------------------

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
