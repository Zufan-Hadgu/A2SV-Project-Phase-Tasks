package data

import (
	"context"
	"log"
	"task_management_api/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongodbURI(mongoStr  string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoStr)

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client,err := mongo.Connect(ctx,clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	return client	
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

func (t *TManager) GetAllTasks() ([]*models.Task,error){
	var tasks []*models.Task
	cur, err := t.Collection.Find(context.TODO(),bson.D{{}})
	if err != nil{
		return nil, err
	}

	for cur.Next(context.TODO()){
		var task  models.Task
		err := cur.Decode(&task)
		if err != nil{
			log.Fatal(err)

		}
		tasks = append(tasks,&task)
	}
	if err := cur.Err() ; err != nil{
		return nil,err
	}
	return tasks,nil

}

func (t *TManager) GetTaskByID(taskID string) (models.Task,error){

	var task models.Task
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return models.Task{},err
	}
	filter := bson.D{{"_id",id}}
	taskErr := t.Collection.FindOne(context.TODO(),filter).Decode(&task)

	if taskErr != nil{
		return models.Task{}, taskErr
	}
	return task,nil

}

func ( t *TManager) AddTask(newTask models.Task) error{
	_,err := t.Collection.InsertOne(context.TODO(),newTask)
	if err != nil{
		return err
	}
	return nil
}

func (t *TManager) UpdateTask(taskID string, updatedTask models.Task) error{
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return err
	}
	filter := bson.D{{"_id", id}}

	update := bson.D{{Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "description", Value: updatedTask.Description},
			{Key: "due_date", Value: updatedTask.DueDate},
			{Key: "status", Value: updatedTask.Status},
		}}}

		_,updateerr := t.Collection.UpdateOne(context.TODO(),filter,update)

		if updateerr != nil{
			return updateerr
		}
		return nil
}

 func (t *TManager) DeteleTask(taskID string) error{
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return err
	}
	filter := bson.D{{"_id", id}}

	_,deleteErr := t.Collection.DeleteOne(context.TODO(),filter)
	if deleteErr != nil{
		return deleteErr
	}
	return nil


 }


