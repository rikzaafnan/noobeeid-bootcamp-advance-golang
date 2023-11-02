package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sesi_7/database"
	"time"
)

func main() {

	timeout := 5 * time.Second
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	defer cancel()

	uri := "mongodb://admin:admin-password@localhost:27017/"
	mongodb, err := database.ConnectMongo(ctx, uri)
	if err != nil {
		log.Println("db not connected with error ", err.Error())
		return
	}

	if mongodb == nil {
		log.Println("db not connected with unknown error ")
		return
	}

	log.Println("db connected")

	Insert(ctx, mongodb, Auth{
		//Id:   primitive.NewObjectID(),
		Name: "Reyhan 2",
	})

}

type Auth struct {
	Id   primitive.ObjectID `bson:"-"`
	Name string             `bson:"name"`
}

func Insert(ctx context.Context, client *mongo.Client, auth Auth) {
	log.Println("try to insert data")

	collection := client.Database("nbid-intermediate").Collection("sesi_7")
	result, err := collection.InsertOne(ctx, auth)
	if err != nil {
		return
	}

	log.Printf("%+v", result)
	log.Println("insert data done")
}
