package Infrastructure

import (
	"context"
	"log"
	"time"

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