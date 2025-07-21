package data

import (
	"context"
	"log"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnetDB(mongostr string) *mongo.Client{

	clientOption := options.Client().ApplyURI(mongostr)

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	
	Client,err := mongo.Connect(ctx,clientOption)
	if err != nil{
		log.Fatal(err)
	}
	return Client


}

type TManager struct{
	Collection *mongo.Collection
}
type TaskManager interface{
	GetAllTasks() ([] *models.Task,error)
	GetTaskByID(taskID string ) (models.Task,error)
	UpdateTask(taskID string, updatedTask models.Task) error
	DeteleTask(taskID string) error
	AddTask(newTask models.Task) error

}

func ( t *TManager) AddTask(newTask models.Task) error{

	_,err := t.Collection.InsertOne(context.TODO(),newTask)
	if err != nil{
		return err
	}
	return nil
}
func (t *TManager) GetAllTasks() ([] *models.Task,error){
	var ListTask []*models.Task
	cur,err := t.Collection.Find(context.TODO(),bson.D{{}}) // curr is aitratior over document, since 

	if err != nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var task models.Task
		if err := cur.Decode(&task); err != nil{
			log.Fatal(err)
		}
		ListTask = append(ListTask,&task)
	}
	if err := cur.Err() ; err != nil{
		return nil,err
	}
	return ListTask,nil

}

func (t *TManager) GetTaskByID(taskID string ) (models.Task,error){
	var task models.Task
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return models.Task{},nil
	}

	filter := bson.D{{"_id",id}}


	TaskErr := t.Collection.FindOne(context.TODO(),filter).Decode(&task)
	if TaskErr != nil{
		return models.Task{},err
	}
	return task,nil

}

func (t *TManager) UpdateTask(taskID string, updatedTask models.Task) error{
	id , err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return err
	}
	filter := bson.D{{"_id",id}}
	update := bson.D{{
		Key:"$set" , Value : bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "description", Value: updatedTask.Description},
			{Key: "due_date", Value: updatedTask.DueDate},
			{Key: "status", Value: updatedTask.Status},
		}}}

		_,updateErr := t.Collection.UpdateOne(context.TODO(),filter,update)
		if updateErr != nil{
			return updateErr

		}

		return nil	 
} 

func ( t *TManager) DeletDeteleTask(taskID string) error{
	id , err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return err
	}
	filter := bson.D{{"_id",id}}
	_,Delerr := t.Collection.DeleteOne(context.TODO(),filter)
	if Delerr != nil{
		return Delerr
	}
	return nil
}


