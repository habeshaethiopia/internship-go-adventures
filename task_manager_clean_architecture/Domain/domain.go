package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	DueDate     time.Time          `json:"due_date"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type TaskUsecase interface {
	CreateTask(task *Task) error
	GetTaskByID(id string) (*Task, error)
	GetTasks() ([]*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
}
type TaskRepository interface {
	CreateTask(task *Task) error
	GetTaskByID(id primitive.ObjectID) (*Task, error)
	GetTasks() ([]*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id primitive.ObjectID) error
}

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `json:"name"`
	Email      string             `json:"email"`
	Password   string             `json:"password"`
	Role       string             `json:"role"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type UserUsecase interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	Login(u User) (User, error)
}
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
	GetUserByEmail(email string) (User, error)
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
