package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sesi_7/database"
)

type RepositoryMongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewRepositoryMongo(client *mongo.Client, db *mongo.Database) (repo RepositoryMongo, err error) {
	if client == nil {
		err = errors.New("client is nil")
		return
	}

	if db == nil {
		err = errors.New("db is nil")
	}

	repo = RepositoryMongo{
		client: client,
		db:     db,
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

func (r RepositoryMongo) FindOneById(ctx context.Context, id string) (model Auth, err error) {

	collection := r.db.Collection(model.GetCollection())
	err = collection.FindOne(ctx, bson.M{"email": id}).Decode(&model)

	return
}
