package usecases

import(
	"github.com/username/taskmanager/domain"
)
type taskUsecase struct {
	taskRepo domain.TaskRepository
}