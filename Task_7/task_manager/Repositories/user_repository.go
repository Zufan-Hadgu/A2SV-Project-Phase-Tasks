package Repositories

import (
	"context"
	"errors"
	"task_manager/Domain"
	"task_manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type UserRepoImpl struct{
	Collection *mongo.Collection

}

func NewUserRepoImpl(col *mongo.Collection) Domain.UserRepository {
	return &UserRepoImpl{Collection: col}
}


func (r *UserRepoImpl) Register(ctx context.Context, user Domain.User) error {
	filter := bson.M{"username": user.Username}
	err := r.Collection.FindOne(ctx, filter).Err()
	if err == nil {
		return errors.New("user already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err 
	}
	hashedPassword, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	_, insertErr := r.Collection.InsertOne(ctx, user)
	return insertErr
}

func (r *UserRepoImpl) Login(ctx context.Context,loginRequest Domain.LoginRequest) (string,error) {
	filter := bson.M{"username":loginRequest.Username}
	var existing Domain.User
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&existing)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
	}
	token,err := Infrastructure.JwtService(existing,loginRequest.Password)
	if err != nil{
		return "",err
	}

	return token , err
	
}
