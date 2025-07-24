package Domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
) 



type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string`bson: "username json:"usrname"`
	Password    string `bson:"password" json:"password"`
	Role        string `bson:"role" json:"role"`
}

type Task struct {
	ID          primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Title       string             `bson: "title" json:"title"`
	Description string             `bson: "description" json:"description"`
	DueDate     time.Time          `bson: "due_date" json:"due_data"`
	Status      string             `bson :"status" json :"status"`
}


type LoginRequest struct {
    Username string `form:"username" json:"username"`
    Password string `form:"password" json:"password"`
}

type TaskRepository interface {
	GetTaskByID(ctx context.Context, taskID string) (Task, error)
	AddTask(ctx context.Context, task Task) error
	UpdateTask(ctx context.Context, taskID string, updatedTask Task) error
	DeleteTask(ctx context.Context,taskID string) error

	GetAllTask(ctx context.Context) ([]Task,error)
}

type UserRepository interface {
	Register(ctx context.Context, user User) error
	Login(ctx context.Context,loginRequest LoginRequest) (string,error)
}