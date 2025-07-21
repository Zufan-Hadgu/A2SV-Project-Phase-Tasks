package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Task struct {
	ID          primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Title       string             `bson: "title" json:"title"`
	Description string             `bson: "description" json:"description"`
	DueDate     time.Time          `bson: "due_date" json:"due_data"`
	Status      string             `bson :"status" json :"status"`
}