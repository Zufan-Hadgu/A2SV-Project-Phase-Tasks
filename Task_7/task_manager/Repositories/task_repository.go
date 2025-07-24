package Repositories

import (
	"context"
	"errors"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepoImpl struct {
	Collection *mongo.Collection
}

func NewTaskRepoImpl(col *mongo.Collection) Domain.TaskRepository {
	return &TaskRepoImpl{Collection: col}
}

func BuildIDFilter (id string)(bson.M,error){
	objID,err := primitive.ObjectIDFromHex(id)

	if err != nil{
		return nil,errors.New("invalid ID format")
	}
	return bson.M{"_id" : objID},nil
}

func (r *TaskRepoImpl) GetTaskByID(ctx context.Context, taskID string) (Domain.Task, error) {
	var task Domain.Task
	filter,err := BuildIDFilter(taskID)
	if err != nil{
		return Domain.Task{} , err

	}
	GetTaskErr := r.Collection.FindOne(ctx, filter).Decode(&task)
	if GetTaskErr != nil{
		return Domain.Task{},GetTaskErr
	}

	return task, nil
}

func ( r *TaskRepoImpl) AddTask(ctx context.Context,task Domain.Task) error{
	_,err:= r.Collection.InsertOne(ctx,task)
	if err != nil{
		return errors.New("Error while creating task")
	}
	return nil
}

func ( r *TaskRepoImpl) UpdateTask(ctx context.Context,taskID string,UpdatedTask Domain.Task) error{
	filter ,err := BuildIDFilter(taskID)
	if err != nil{
		return err

	}
	update := bson.D{{
		Key:"$set" , Value : bson.D{
			{Key: "title", Value: UpdatedTask.Title},
			{Key: "description", Value: UpdatedTask.Description},
			{Key: "due_date", Value: UpdatedTask.DueDate},
			{Key: "status", Value: UpdatedTask.Status},
		}}}

		_,updateErr := r.Collection.UpdateOne(context.TODO(),filter,update)
		if updateErr != nil{
			return updateErr

		}

		return nil	
}
func ( r *TaskRepoImpl) DeleteTask(ctx context.Context,taskID string) error{
	filter ,err := BuildIDFilter(taskID)
	if err != nil{
		return err

	}
   _,delErr:= r.Collection.DeleteOne(ctx,filter)
   if delErr != nil{
	return delErr
   }
   return nil
}

func ( r *TaskRepoImpl) GetAllTask(ctx context.Context) ([]Domain.Task,error){
	var ListOfTasks []Domain.Task
	task,err := r.Collection.Find(ctx,bson.D{{}})
	if err != nil{
		return nil,err
	}

	for task.Next(ctx){
		var Ntask Domain.Task
		if err := task.Decode(&Ntask); err != nil{
			return nil,err
		}
		ListOfTasks = append(ListOfTasks,Ntask)

	}
	return ListOfTasks,nil

}