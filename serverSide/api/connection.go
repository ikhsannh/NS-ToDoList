package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://nashkispace:b6Ol0nK2DPZRcXPQ@ac-sfoy2es-shard-00-00.clkur4r.mongodb.net:27017,ac-sfoy2es-shard-00-01.clkur4r.mongodb.net:27017,ac-sfoy2es-shard-00-02.clkur4r.mongodb.net:27017/todolist?ssl=true&replicaSet=atlas-slxwf6-shard-0&authSource=admin&retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("todolist").Collection(collectionName)
	return collection
}
