package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string               `json:"email"`
	Password string               `json:"password"`
	Role     string               `json:"role"`
	TaskIDs  []primitive.ObjectID `bson:"task_ids,omitempty" json:"task_ids,omitempty"`
}
