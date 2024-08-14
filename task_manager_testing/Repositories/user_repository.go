package repositories

import (
	"context"
	"fmt"

	domain "task/Domain"

	"task/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// GetUserByEmail implements domain.UserRepository.

// GetTaskByEmail implements domain.UserRepository.

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (u *userRepository) GetUserByEmail(email string) (domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(user *domain.User) error {

	collection := u.database.Collection(u.collection)
	var user1 domain.User
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&user1)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	_, err = collection.InsertOne(context.Background(), user)

	return err

}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(id primitive.ObjectID) error {
	collection := u.database.Collection(u.collection)

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult)
	return nil
}

// GetUserByID implements domain.UserRepository.
func (u *userRepository) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil

}

// GetUsers implements domain.UserRepository.
func (u *userRepository) GetUsers() ([]*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var users []*domain.User
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user *domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(user *domain.User) error {
	collection := u.database.Collection(u.collection)

	updateResult, err := collection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v documents in the trainers collection\n", updateResult.ModifiedCount)
	return nil
}
