package data

import (
	"context"
	"fmt"
	"log"
	"task/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users *mongo.Collection

func FindAllUsers(collection *mongo.Collection) []*models.User {
	findOptions := options.Find()

	var results []*models.User
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.User
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

func FindOneUser(collection *mongo.Collection, filter bson.M) (models.User, error) {

	var result models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.User{}, nil
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, nil
}
func DeleteUser(collection *mongo.Collection, filter bson.M) error {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return nil
}
func InsertUser(collection *mongo.Collection, data models.User) error {

	// insert data
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func UpdateUser(collection *mongo.Collection, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult, nil
}