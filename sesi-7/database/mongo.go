package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.context, uri string) (client *mongo.Client, err error) {

	opts := options.Client().ApplyURI(uri)

	return
}
