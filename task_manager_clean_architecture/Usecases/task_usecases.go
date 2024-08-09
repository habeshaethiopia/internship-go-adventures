package usecases

import (
	domain "task/Domain"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
// CreateTask implements domain.TaskUsecase.
func (t *taskUsecase) CreateTask(task *domain.Task) error {
	
	panic("unimplemented")
}

// DeleteTask implements domain.TaskUsecase.
func (t *taskUsecase) DeleteTask(id string) error {

	panic("unimplemented")
}

// GetTaskByID implements domain.TaskUsecase.
func (t *taskUsecase) GetTaskByID(id string) (*domain.Task, error) {
	panic("unimplemented")
}

// GetTasks implements domain.TaskUsecase.
func (t *taskUsecase) GetTasks() ([]*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements domain.TaskUsecase.
func (t *taskUsecase) UpdateTask(task *domain.Task) error {
	panic("unimplemented")
}

