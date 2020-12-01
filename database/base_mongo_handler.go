package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BaseMongoHandler struct {
	collection *mongo.Collection
}

func (bmh *BaseMongoHandler) getBsonIDorPanic(errorKey string, id string) primitive.ObjectID {
	bsonID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panicf("invalid data: %s", errorKey)
	}
	return bsonID
}

func (bmh *BaseMongoHandler) must(args ...interface{}) interface{} {
	v, err := args[0], args[1]
	if err != nil {
		panic(err)
	}

	return v
}
