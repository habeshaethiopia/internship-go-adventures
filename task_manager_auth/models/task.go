package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	DueDate     time.Time          `bson:"due_date"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}
