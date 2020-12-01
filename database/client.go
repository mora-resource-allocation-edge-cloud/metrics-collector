package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"metrics-collector/conf"
	"strings"
)

var VideoReproductionRepository VideoReproductionHandler

func init() {
	log.Println("Initializing MongoDB database connector")
	var args []string
	if len(conf.MongoAuthSource) > 0 {
		args = append(args, fmt.Sprintf("authSource=%s", conf.MongoAuthSource))
	}
	argStr := strings.Join(args, "&")
	if len(argStr) > 0 {
		argStr = "?" + argStr
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d/%s%s",
			conf.MongoUser, conf.MongoPassword, conf.MongoHost, conf.MongoPort, conf.MongoDBName, argStr,
		)))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dbHandler := client.Database(conf.MongoDBName)

	mongoHandler := videoReproductionMongoHandler{
		BaseMongoHandler{
			collection: dbHandler.Collection("metrics"),
		},
	}
	VideoReproductionRepository = &mongoHandler
	log.Println("MongoDB connection has been initialized")
}
