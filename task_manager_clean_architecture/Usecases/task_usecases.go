package usecases

import (
	"fmt"
	domain "task/Domain"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

// CreateTask implements domain.TaskUsecase.
func (t *taskUsecase) CreateTask(task *domain.Task) error {
	task.ID = primitive.NewObjectID()
	color.Red("usecase create task", task)
	return t.taskRepository.CreateTask(task)

}

// DeleteTask implements domain.TaskUsecase.
func (t *taskUsecase) DeleteTask(id string) error {
	idx, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	color.Red("usecase delete task")
	return t.taskRepository.DeleteTask(idx)

}

// GetTaskByID implements domain.TaskUsecase.
func (t *taskUsecase) GetTaskByID(id string) (*domain.Task, error) {
	idx, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.Task{}, err
	}
	color.Red("usecase get task by id")
	return t.taskRepository.GetTaskByID(idx)
}

// GetTasks implements domain.TaskUsecase.
func (t *taskUsecase) GetTasks() ([]*domain.Task, error) {

	return t.taskRepository.GetTasks()
}

// UpdateTask implements domain.TaskUsecase.
// UpdateTask implements domain.TaskUsecase.
func (t *taskUsecase) UpdateTask(task *domain.Task) error {
	existingTask, err := t.taskRepository.GetTaskByID(task.ID)
	
	if err != nil {
		fmt.Print(existingTask)
		return err
	}

	if task.Title == "" {
		task.Title = existingTask.Title
	}
	if task.Description == "" {
		task.Description = existingTask.Description
	}
	if task.Status == "" {
		task.Status = existingTask.Status
	}
	task.UserID = existingTask.UserID
	task.DueDate = existingTask.DueDate
	color.Red("usecase update task")

	return t.taskRepository.UpdateTask(task)
}
