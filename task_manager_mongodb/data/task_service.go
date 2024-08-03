// save the tasks to json file
package data

import (
	"context"
	"fmt"
	"log"

	"task/models"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "time"
)

func CreateDB() (*mongo.Collection, *mongo.Client) {

	// Create a MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// connect to mongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("Tasks").Collection("task")

	return collection, client
}

var Tasks *mongo.Collection

func FindAll(collection *mongo.Collection) []*models.Task {
	findOptions := options.Find()

	var results []*models.Task
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
		fmt.Println(elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}
func FindOne(collection *mongo.Collection, filter bson.M) (models.Task, error) {
    
	var result models.Task
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, err
		}
		return models.Task{}, err
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, nil
}
func Delete(collection *mongo.Collection, filter bson.M) error {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return nil
}
func InsertOne(collection *mongo.Collection, data models.Task) error {

	// insert data
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// var abebe models.Task {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"}
func UpdateOne(collection *mongo.Collection, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult, nil
}
func CloseDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
