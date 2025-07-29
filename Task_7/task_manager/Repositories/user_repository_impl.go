package Repositories

import (
	"context"
	"errors"
	"log"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type UserRepoImpl struct{
	Collection *mongo.Collection

}

func NewUserRepoImpl(col *mongo.Collection) Domain.IUserRepository {
	return &UserRepoImpl{Collection: col}
}

func (r *UserRepoImpl) FindByUsername(username string) (*Domain.User,error){
	var user Domain.User
	filter := bson.M{"username":username}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return nil , errors.New("No user found")
		}
		return nil,errors.New("Database error")
	}
	return &user,nil
}



func (r *UserRepoImpl) Create(user Domain.User) error {
	_, insertErr := r.Collection.InsertOne(context.TODO(), user)
	log.Println(insertErr)
	return insertErr
}

func (r *UserRepoImpl) CountDB()(int64,error){
	count, err := r.Collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return 0,errors.New("databse error")
	}
	return count,nil

}

