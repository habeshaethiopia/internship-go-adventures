package infrastructure

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "time"
)

func ConnectDB(url string, DBname string) (*mongo.Database, *mongo.Client, error) {
	// Create a MongoDB client
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil,nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil,err
	}

	fmt.Println("Connected to MongoDB!")
	DB:= client.Database(DBname)
	return DB,client, nil
}

func CreateCollection(database mongo.Database,  collectionName string) *mongo.Collection {
	collection := database.Collection(collectionName)
	return collection
}

func CloseDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}

