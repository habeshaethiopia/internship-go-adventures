package usecases

import (
	"errors"
	domain "task/Domain"
	mocks "task/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecaseSuite defines the suite for task usecase tests
type TaskUseCaseSuite struct {
	suite.Suite
	taskUseCase  domain.TaskUsecase   // Use case under test
	mockTaskRepo mocks.TaskRepository // Mocked repository
}

// SetupTest sets up the test environment
func (suite *TaskUseCaseSuite) SetupTest() {
	suite.taskUseCase = NewTaskUsecase(&suite.mockTaskRepo)
}

// TestExample is an example test case
func (suite *TaskUseCaseSuite) TestCreateTask() {
	suite.mockTaskRepo.On("CreateTask", mock.Anything).Return(nil).Once() // Mocking repository behavior
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}

	err := suite.taskUseCase.CreateTask(&demotask)

	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "CreateTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestCreateTaskError() {
	suite.mockTaskRepo.On("CreateTask", mock.Anything).Return(nil).Once() // Mocking repository behavior
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}

	err := suite.taskUseCase.CreateTask(&demotask)

	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "CreateTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestCreateTaskErrorRepo() {
	suite.mockTaskRepo.On("CreateTask", mock.Anything).Return(nil).Once() // Mocking repository behavior
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}

	err := suite.taskUseCase.CreateTask(&demotask)

	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "CreateTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestCreateTaskErrorRepo2() {
	suite.mockTaskRepo.On("CreateTask", mock.Anything).Return(nil).Once() // Mocking repository behavior
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}

	err := suite.taskUseCase.CreateTask(&demotask)

	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "CreateTask", mock.Anything)
}

func (suite *TaskUseCaseSuite) TestUpdateTask() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(&domain.Task{}, nil).Once()
	suite.mockTaskRepo.On("UpdateTask", mock.Anything).Return(nil)
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}
	err := suite.taskUseCase.UpdateTask(&demotask)
	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
	suite.mockTaskRepo.AssertCalled(suite.T(), "UpdateTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestUpdateTaskNotFound() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(nil, errors.New("some error")).Once()
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		Status:      "progress",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}
	err := suite.taskUseCase.UpdateTask(&demotask)
	suite.Error(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestUpdateTaskError() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(&domain.Task{}, nil).Once()
	suite.mockTaskRepo.On("UpdateTask", mock.Anything).Return(nil)
	demotask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now(),
		UserID:      primitive.NewObjectID(),
	}
	err := suite.taskUseCase.UpdateTask(&demotask)
	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
	suite.mockTaskRepo.AssertCalled(suite.T(), "UpdateTask", mock.Anything)
}

func (suite *TaskUseCaseSuite) TestDeleteTask() {
	suite.mockTaskRepo.On("DeleteTask", mock.Anything).Return(nil).Once()
	err := suite.taskUseCase.DeleteTask("5f4c6d1b6b4b3f3f4c7f3f4c")
	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "DeleteTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestDeleteTaskError() {
	suite.mockTaskRepo.On("DeleteTask", mock.Anything).Return(errors.New("error when deleteing the task")).Once() // Mocking repository behavior to return an error
	id := primitive.NewObjectID()
	err := suite.taskUseCase.DeleteTask(id.Hex())
	suite.Error(err) // Expecting an error
	suite.mockTaskRepo.AssertCalled(suite.T(), "DeleteTask", mock.AnythingOfType("primitive.ObjectID"))
}
func (suite *TaskUseCaseSuite) TestDeleteTaskErrorNotFound() {
	suite.mockTaskRepo.On("DeleteTask", mock.Anything).Return(errors.New("task not found")).Once()
	err := suite.taskUseCase.DeleteTask(primitive.NewObjectID().Hex())
	suite.Error(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "DeleteTask", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestGetTaskByID() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(&domain.Task{}, nil).Once()

	_, err := suite.taskUseCase.GetTaskByID(primitive.NewObjectID().Hex())
	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestGetTaskByIDError() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(nil, errors.New("error")).Once()
	_, err := suite.taskUseCase.GetTaskByID(primitive.NewObjectID().Hex())
	suite.Error(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestGetTaskByIDErrorNotFound() {
	suite.mockTaskRepo.On("GetTaskByID", mock.Anything).Return(nil, errors.New("task not found")).Once()
	_, err := suite.taskUseCase.GetTaskByID(primitive.NewObjectID().Hex())
	suite.Error(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
}
func (suite *TaskUseCaseSuite) TestGetTasks() {
	suite.mockTaskRepo.On("GetTasks").Return([]*domain.Task{}, nil).Once()
	_, err := suite.taskUseCase.GetTasks()
	suite.NoError(err)
	suite.mockTaskRepo.AssertCalled(suite.T(), "GetTasks")
}

// TestTaskUsecaseSuite runs the test suite
func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}
