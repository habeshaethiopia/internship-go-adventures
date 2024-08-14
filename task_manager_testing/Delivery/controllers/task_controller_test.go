package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	domain "task/Domain"
	mocks "task/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerTestSuite struct {
	suite.Suite
	taskController  *TaskController    // Use case under test
	mockTaskUsecase *mocks.TaskUsecase // Mocked repository
	testserver      *httptest.Server
}

func (suite *TaskControllerTestSuite) SetupTest() {

	R := gin.Default()
	R.GET("/tasks", suite.taskController.GetTasks)
	R.GET("/tasks/:id", suite.taskController.GetTaskByID)
	R.POST("/task", suite.taskController.CreateTask)
	R.PUT("/tasks/:id", suite.taskController.UpdateTask)
	R.DELETE("/tasks/:id", suite.taskController.DeleteTask)

	suite.testserver = httptest.NewServer(R)
	suite.mockTaskUsecase = new(mocks.TaskUsecase)
	suite.taskController = NewTaskController(suite.mockTaskUsecase)
}
func (suite *TaskControllerTestSuite) TearDownSuite() {
	suite.testserver.Close()
}
func (suite *TaskControllerTestSuite) TestCreateTask_Success() {
	// Prepare the request and context
	reqBody := `{"title":"test","description":"test","status":"test"}`
	req, err := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(reqBody))
	if err != nil {
		suite.FailNow("error when creating the request")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	id := primitive.NewObjectID().Hex()
	c.Set("claims", &domain.Claims{UserID: id}) // Example ObjectID

	// Prepare the expected task object and set up mock behavior
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		suite.FailNow("premitive object id is failed")
	}
	expectedTask := &domain.Task{
		Title:       "test",
		Description: "test",
		Status:      "test",
		UserID:      userID,
	}

	suite.mockTaskUsecase.On("CreateTask", expectedTask).Return(nil).Once()

	// Create the task controller and call the CreateTask method
	suite.taskController.CreateTask(c)

	// Verify the response
	if w.Code != http.StatusCreated {
		suite.T().Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
	}

	var responseTask domain.Task
	err = json.NewDecoder(w.Body).Decode(&responseTask)
	if err != nil {
		suite.T().Fatal(err)
	}
	if !reflect.DeepEqual(responseTask, *expectedTask) {
		suite.T().Errorf("expected response body %+v but got %+v", *expectedTask, responseTask)
	}

	suite.mockTaskUsecase.AssertCalled(suite.T(), "CreateTask", mock.AnythingOfType("*domain.Task"))
}
func (suite *TaskControllerTestSuite) TestDeleteTask_Success() {
	// Prepare the request and context
	req, err := http.NewRequest(http.MethodDelete, "/tasks/123", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	id := primitive.NewObjectID().Hex()
	c.Set("claims", &domain.Claims{UserID: id}) // Example ObjectID
	// Prepare the expected task object and set up mock behavior
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		suite.T().Fatal(err)
	}
	expectedTask := &domain.Task{
		ID:     primitive.NewObjectID(),
		Title:  "test",
		Status: "test",
		UserID: userID,
	}
	suite.mockTaskUsecase.On("GetTaskByID", mock.Anything).Return(expectedTask, nil).Once()
	suite.mockTaskUsecase.On("DeleteTask", mock.Anything).Return(nil).Once()

	// Create the task controller and call the DeleteTask method
	suite.taskController.DeleteTask(c)

	// Verify the response
	if w.Code != http.StatusOK {
		suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var response gin.H
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	expectedResponse := gin.H{"message": "Task deleted successfully"}
	if !reflect.DeepEqual(response, expectedResponse) {
		suite.T().Errorf("expected response body %+v but got %+v", expectedResponse, response)
	}

	suite.mockTaskUsecase.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
	suite.mockTaskUsecase.AssertCalled(suite.T(), "DeleteTask", mock.Anything)
}
func (suite *TaskControllerTestSuite) TestGetTaskByID() {
	// Prepare the request and context
	suite.Run("Success", func() {
		req, err := http.NewRequest(http.MethodGet, "/tasks/123", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID
		// Prepare the expected task object and set up mock behavior
		userID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			suite.T().Fatal(err)
		}
		expectedTask := &domain.Task{
			ID:     primitive.NewObjectID(),
			Title:  "test",
			Status: "test",
			UserID: userID,
		}
		suite.mockTaskUsecase.On("GetTaskByID", mock.Anything).Return(expectedTask, nil).Once()

		// Create the task controller and call the GetTaskByID method
		suite.taskController.GetTaskByID(c)

		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		var response domain.Task
		err = json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(response, *expectedTask) {
			suite.T().Errorf("expected response body %+v but got %+v", *expectedTask, response)
		}

		suite.mockTaskUsecase.AssertCalled(suite.T(), "GetTaskByID", mock.Anything)
	})
}
func (suite *TaskControllerTestSuite) TestGetTasks_Success() {
	// Prepare the request and context
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	id := primitive.NewObjectID().Hex()
	c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID
	// Prepare the expected task object and set up mock behavior
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		suite.T().Fatal(err)
	}
	expectedTasks := []*domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "test",
			Description: "",
			Status:      "progressed",

			UserID: userID,
		},
	}
	suite.mockTaskUsecase.On("GetTasks").Return(expectedTasks, nil).Once()

	// Create the task controller and call the GetTasks method
	suite.taskController.GetTasks(c)

	// Verify the response
	if w.Code != http.StatusOK {
		suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var response []domain.Task
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	expectedresult := []domain.Task{
		{
			ID:          expectedTasks[0].ID,
			Title:       "test",
			Description: "",
			Status:      "progressed",
			DueDate:     expectedTasks[0].DueDate,
			UserID:      expectedTasks[0].UserID},
	}
	suite.Equal(expectedresult, response)
	// if !reflect.DeepEqual(response, expectedTasks) {
	// 	suite.T().Errorf("expected response body %+v but got %+v", expectedresult, response)
	// }

	suite.mockTaskUsecase.AssertCalled(suite.T(), "GetTasks")
}
func (suite *TaskControllerTestSuite) TestUpdateTask() {
	suite.Run("success", func () {
		// Prepare the request and context
		reqBody := `{"title":"test","description":"test","status":"test"}`
		id := primitive.NewObjectID().Hex()
		req, err := http.NewRequest(http.MethodPut, "/tasks/"+id, strings.NewReader(reqBody))
		if err != nil {
			suite.T().Fatal(err)
		}
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: id}}
	
		// Prepare the expected task object and set up mock behavior
		userID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			suite.T().Fatal(err)
		}
		expectedTask := &domain.Task{
			ID:          userID,
			Title:       "test",
			Description: "test",
			Status:      "test",
		
		}
		suite.mockTaskUsecase.On("UpdateTask", expectedTask).Return(nil).Once()
	
		// Create the task controller and call the UpdateTask method
		suite.taskController.UpdateTask(c)
	
		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		var response domain.Task
		err = json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(response, *expectedTask) {
			suite.T().Errorf("expected response body %+v but got %+v", *expectedTask, response)
		}
	
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("error", func () {
		// Prepare the request and context
		reqBody := `{"title":"test","description":"test","status":"test"}`
		id := primitive.NewObjectID().Hex()
		req, err := http.NewRequest(http.MethodPut, "/tasks/"+id, strings.NewReader(reqBody))
		if err != nil {
			suite.T().Fatal(err)
		}
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: id}}
	
		// Prepare the expected task object and set up mock behavior
		userID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			suite.T().Fatal(err)
		}
		expectedTask := &domain.Task{
			ID:          userID,
			Title:       "test",
			Description: "test",
			Status:      "test",
		
		}
		suite.mockTaskUsecase.On("UpdateTask", expectedTask).Return(nil).Once()
	
		// Create the task controller and call the UpdateTask method
		suite.taskController.UpdateTask(c)
	
		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		var response domain.Task
		err = json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(response, *expectedTask) {
			suite.T().Errorf("expected response body %+v but got %+v", *expectedTask, response)
		}
	
		suite.Equal(http.StatusOK, w.Code)
	})
	
}
func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
