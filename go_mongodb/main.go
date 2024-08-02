package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

var ash = Trainer{"Ash", 10, "Pallet Town"}
var misty = Trainer{"Misty", 10, "Cerulean City"}
var brock = Trainer{"Brock", 15, "Pewter City"}

func InsertMany(collection *mongo.Collection) {
	fmt.Println("Inserted a single document: ", ash)
	// insert multiple data
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.Background(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
func InsertOne(collection *mongo.Collection, data Trainer) {

	// insert data
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateOne(collection *mongo.Collection, filter bson.D, update bson.D) {
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func UpdateMany(collection *mongo.Collection, filter bson.D, update bson.D) {
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func Find( collection *mongo.Collection) {
	// Pass these options to the Find method
findOptions := options.Find()
// findOptions.SetLimit(2)

// Here's an array in which you can store the decoded documents
var results []*Trainer

// Passing bson.D{{}} as the filter matches all documents in the collection
cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem Trainer
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

// Close the cursor once finished
cur.Close(context.TODO())

fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

}
func Delete(collection *mongo.Collection) {
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{Key: "name", Value: "Ash"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
func main() {
	// set client option
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

	collection := client.Database("test").Collection("trainers")
	// InsertOne(collection, ash)
	// InsertMany(collection)

	//Update train
	// filter := bson.D{{Key: "name", Value: "Ash"}}
	// update := bson.D{
	// 	{Key: "$inc", Value: bson.D{
	// 		{Key: "age", Value: 1},
	// 	}},
	// }
	// update one data
	// UpdateOne(collection,filter,update)
	// UpdateMany(collection,filter,update)

	//find data 
	Find(collection)
	//delete data
	Delete(collection)
	Find(collection)
	// disconnect from mongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
