package repositories

import (
	domain "task/Domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

// CreateTask implements domain.TaskRepository.
func (t *taskRepository) CreateTask(task *domain.Task) error {
	panic("unimplemented")
}

// DeleteTask implements domain.TaskRepository.
func (t *taskRepository) DeleteTask(id string) error {
	panic("unimplemented")
}

// GetTaskByID implements domain.TaskRepository.
func (t *taskRepository) GetTaskByID(id string) (*domain.Task, error) {
	panic("unimplemented")
}

// GetTasks implements domain.TaskRepository.
func (t *taskRepository) GetTasks() ([]*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements domain.TaskRepository.
func (t *taskRepository) UpdateTask(task *domain.Task) error {
	panic("unimplemented")
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}
