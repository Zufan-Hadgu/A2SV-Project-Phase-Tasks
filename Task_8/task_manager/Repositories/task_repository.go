package Repositories

import (
	"errors"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

type TaskRepoImpl struct {
	Collection *mongo.Collection
}

func NewTaskRepoImpl(col *mongo.Collection) Domain.ITaskRepository {
	return &TaskRepoImpl{Collection: col}
}

func BuildIDFilter(id string) (bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	return bson.M{"_id": objID}, nil
}

func (r *TaskRepoImpl) GetTaskByID(taskID string) (Domain.Task, error) {
	var task Domain.Task
	filter, err := BuildIDFilter(taskID)
	if err != nil {
		return Domain.Task{}, err
	}
	err = r.Collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return Domain.Task{}, err
	}
	return task, nil
}

func (r *TaskRepoImpl) AddTask(task Domain.Task) error {
	_, err := r.Collection.InsertOne(context.TODO(), task)
	if err != nil {
		return errors.New("error while creating task")
	}
	return nil
}

func (r *TaskRepoImpl) UpdateTask(taskID string, UpdatedTask Domain.Task) error {
	filter, err := BuildIDFilter(taskID)
	if err != nil {
		return err
	}

	update := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "title", Value: UpdatedTask.Title},
			{Key: "description", Value: UpdatedTask.Description},
			{Key: "due_date", Value: UpdatedTask.DueDate},
			{Key: "status", Value: UpdatedTask.Status},
		},
	}}

	_, err = r.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepoImpl) DeleteTask(taskID string) error {
	filter, err := BuildIDFilter(taskID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepoImpl) GetAllTask() ([]Domain.Task, error) {
	var tasks []Domain.Task
	cursor, err := r.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
