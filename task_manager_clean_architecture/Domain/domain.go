package domain
import (
	"time"

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
	GetTaskByID(id string) (*Task, error)
	GetTasks() ([]*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
}


type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string               `json:"email"`
	Password string               `json:"password"`
	Role     string               `json:"role"`
}

type UserUsecase interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
		