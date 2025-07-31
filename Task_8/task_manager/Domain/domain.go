package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	DueDate     time.Time `bson:"due_date" json:"due_date"`
	Status      string    `bson:"status" json:"status"`
}


type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}



type ITaskRepository interface {
	GetTaskByID(taskID string) (Task, error)
	AddTask(task Task) error
	UpdateTask(taskID string, updatedTask Task) error
	DeleteTask(taskID string) error
	GetAllTask() ([]Task, error)
}

type IUserRepository interface {
	Create(user User) error
	FindByUsername(username string)(*User,error) 
	CountDB()(int64,error)
}


type ITaskUsecase interface {
	GetTaskByID(taskID string) (Task, error)
	AddTask(task Task) error
	UpdateTask(taskID string, updatedTask Task) error
	DeleteTask(taskID string) error
	GetAllTask() ([]Task, error)
}

type IUserUsecase interface {
	Register(user User)(string,error)
	Login(loginrequest LoginRequest)(string,error)
}

type IJwtService interface {
	GenerateToken(userID string, role string) (string,error)
}


type IPasswordService interface {
	HashPassword (password string) (string,error)
	ComparePassword(hashedPassword,password string,) error 
}