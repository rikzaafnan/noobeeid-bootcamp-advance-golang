package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sesi_7/app/auth"
	"sesi_7/database"
	"time"
)

func main() {

	timeout := 5 * time.Second

	log.Println(timeout)

	//	mongodb
	//mongoDB(timeout)

	//	redis
	redis(timeout)

}

func mongoDB(timeout time.Duration) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	defer cancel()

	uri := "mongodb://admin:admin-password@localhost:27017/"
	mongoClient, err := database.ConnectMongo(ctx, uri)
	if err != nil {
		log.Println("db not connected with error ", err.Error())
		return
	}

	if mongoClient == nil {
		log.Println("db not connected with unknown error ")
		return
	}

	mongoDB := mongoClient.Database(database.DB_NAME)

	log.Println("db connected")

	repo, err := auth.NewRepositoryMongo(mongoClient, mongoDB)
	if err != nil {
		log.Println("error when to try to execute NewRepositoryMongo with detail : ", err.Error())
		return
	}

	//lastId, err := repo.Insert(ctx, auth.Auth{
	//	Email:     "reyhan@mail.com",
	//	CreatedAt: time.Now(),
	//	IsActive:  false,
	//})
	//
	//if err != nil {
	//	log.Println("error when to try to execute repo.insert with detail : ", err.Error())
	//	return
	//}
	//
	//log.Println("insert success with id : ", lastId)
	//Insert(ctx, mongoClient, Auth{
	//	//Id:   primitive.NewObjectID(),
	//	Name: "Reyhan 2",
	//})

	//find one
	res, err := repo.FindOneById(ctx, "reyhan@mail.com")
	if err != nil {
		log.Println("error when to try to execute FindOneById with detail : ", err.Error())
		return
	}

	log.Println("select find one success with id : ", res)
}

func redis(timeout time.Duration) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	defer cancel()
	redisClient, err := database.ConnectRedis(ctx, "localhost:6379", "user_password")
	if err != nil {
		log.Println("db not connected with error ", err.Error())
		return
	}

	if redisClient == nil {
		log.Println("db not connected with unknown error ")
		return
	}
	log.Println("redis connected")

	err = redisClient.Set(ctx, "token-user01", "ini token user contoh", 10*time.Second).Err()
	if err != nil {
		log.Println("error set data ", err.Error())
		return
	}

	log.Println("success create data")

	cmd := redisClient.Get(ctx, "token-user02")
	res, err := cmd.Result()
	if err != nil {
		log.Println("error get data ", err.Error())
		return
	}

	log.Println("isi token token-user02 adalah : ", res)

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
