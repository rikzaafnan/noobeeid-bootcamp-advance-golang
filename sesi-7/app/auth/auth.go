package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Auth struct {
	Id        primitive.ObjectID `bson:"-"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	IsActive  bool               `bson:"is_active"`
}

func (a Auth) GetCollection() (collection string) {
	return "auth"
}
