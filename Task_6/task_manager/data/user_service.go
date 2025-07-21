package data

import (
	"context"
	"errors"
	"log"
	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserManager interface {
	Register(user models.User) error
	Login(username string, password string) error
}

type ImplUser struct {
	Collection mongo.Collection
}

func (u *ImplUser) Register(user models.User) error {

	filter := bson.M{"username": user.Username}
	var existing models.User
	errDB := u.Collection.FindOne(context.TODO(), filter).Decode((&existing))
	if errDB == nil {
		return errors.New("user already exist")
	}
	count, err := u.Collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return errors.New("databse error")
	}
	user.Role = "user"
	if count == 0 {
		user.Role = "admin"
	}

	log.Println(user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	_, err = u.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *ImplUser) Login(username string, password string) (string, error) {
	var jwtSecret = []byte("my-secret")
	filter := bson.M{"username": username}
	var existing models.User
	err := u.Collection.FindOne(context.TODO(), filter).Decode(&existing)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   existing.ID,
		"role": existing.Role,
	})
	token, err := tokenString.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("internal server error")

	}
	return token, nil

}
