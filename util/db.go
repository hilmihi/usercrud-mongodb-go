package util

import (
	"context"
	"fmt"
	"go-mongodb/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoDBConnect(config *config.AppConfig) *mongo.Client {
	fmt.Println(config.Url)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Url))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)

	return client
}
