package dbs

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var mongoErr error
var mongoOnce sync.Once

const (
	CONNECTION_STRING = "mongodb://localhost:27017"
	DB                = "practice_db"
)

func getMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTION_STRING))
		if err != nil {
			log.Fatal(err)
			mongoErr = err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
			mongoErr = err
		}
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
			mongoErr = err
		}
		mongoClient = client
	})
	return mongoClient, mongoErr
}

func GetMongoCollection(collectionName string) *mongo.Collection {
	client, err := getMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(DB).Collection(collectionName)
}
