package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sesi_7/database"
)

type RepositoryMongo struct {
	client *mongo.Client
}

func NewRepositoryMongo(client *mongo.Client) (repo RepositoryMongo, err error) {
	if client == nil {
		err = errors.New("client is nil")
		return
	}

	repo = RepositoryMongo{
		client: client,
	}

	return
}

func (r RepositoryMongo) Insert(ctx context.Context, model Auth) (insertId primitive.ObjectID, err error) {

	collection := r.client.Database(database.DB_NAME).Collection(model.GetCollection())
	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return
	}

	insertID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		err = errors.New("insertId is not primitive.ObjectID")
	}

	return insertID, nil
}
