package repositories

import (
	"context"
	"fmt"
	domain "task/Domain"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task/mongo"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

// CreateTask implements domain.TaskRepository.
func (t *taskRepository) CreateTask(task *domain.Task) error {
	collection := t.database.Collection(t.collection)
	_, err := collection.InsertOne(context.Background(), task)
	return err
}

// DeleteTask implements domain.TaskRepository.
func (t *taskRepository) DeleteTask(id primitive.ObjectID) error {
	collection := t.database.Collection(t.collection)

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult)
	return nil
}

// GetTaskByID implements domain.TaskRepository.
func (t *taskRepository) GetTaskByID(id primitive.ObjectID) (*domain.Task, error) {
	collection := t.database.Collection(t.collection)
	
	var task domain.Task
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return &domain.Task{}, err
	}
	color.Red("repo get task by id")
	return &task, nil
}

// GetTasks implements domain.TaskRepository.
func (t *taskRepository) GetTasks() ([]*domain.Task, error) {
	collection := t.database.Collection(t.collection)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []*domain.Task
	for cursor.Next(context.Background()) {
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

// UpdateTask implements domain.TaskRepository.
func (t *taskRepository) UpdateTask(task *domain.Task) error {
	collection := t.database.Collection(t.collection)
	updateResult, err := collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v documents in the trainers collection\n", updateResult.ModifiedCount)
	return nil
}
