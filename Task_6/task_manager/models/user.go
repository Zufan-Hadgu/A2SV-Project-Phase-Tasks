package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string`bson: "username json:"usrname"`
	Password    string `bson:"password" json:"password"`
	Role        string `bson:"role" json:"role"`
}