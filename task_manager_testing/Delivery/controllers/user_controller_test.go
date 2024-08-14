package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	domain "task/Domain"
	infrastructure "task/Infrastructure"
	mocks "task/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	controller *UserController
	usecase    *mocks.UserUsecase
	router     *gin.Engine
}

func (suite *UserControllerSuite) SetupTest() {
	suite.usecase = new(mocks.UserUsecase)
	suite.controller = NewUserController(suite.usecase)
	suite.router = gin.Default()
	suite.router.POST("/users", suite.controller.CreateUser)
	suite.router.DELETE("/users/:id", suite.controller.DeleteUser)
	suite.router.GET("/users/:id", suite.controller.GetUserByID)
	suite.router.GET("/users", suite.controller.GetUsers)
	suite.router.PUT("/users/:id", suite.controller.UpdateUser)
	suite.router.POST("/users/login", suite.controller.LoginUser)
}
func (suite *UserControllerSuite) TestCreateUser() {
	suite.Run("success", func() {
		reqBody := `{"name":"testuser","email":"test@example.com", "password":"test","role":"user"}`
		req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
		if err != nil {
			suite.FailNow("error when creating the request")
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Prepare the expected user object and set up mock behavior

		expectedUser := &domain.User{
			Name:     "testuser",
			Email:    "test@example.com",
			Password: "test",
			Role:     "user",
		}

		suite.usecase.On("CreateUser", expectedUser).Return(nil).Once()

		// Create the user controller and call the CreateUser method
		suite.controller.CreateUser(c)

		// Verify the response
		if w.Code != http.StatusCreated {
			suite.T().Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
		}

		var responseUser domain.User
		err = json.NewDecoder(w.Body).Decode(&responseUser)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(responseUser.ID, (*expectedUser).ID) {
			suite.T().Errorf("expected response body %+v but got %+v", *expectedUser, responseUser)
		}

		suite.usecase.AssertCalled(suite.T(), "CreateUser", mock.AnythingOfType("*domain.User"))
	})
	suite.Run("error", func() {
		reqBody := `{"name":"testuser","email":"test@example.com", "password":"test","role":"user"}`
		req, err := http.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
		if err != nil {
			suite.FailNow("error when creating the request")
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Prepare the expected user object and set up mock behavior

		expectedUser := &domain.User{
			Name:     "testuser",
			Email:    "test@example.com",
			Password: "test",
			Role:     "user",
		}

		suite.usecase.On("CreateUser", expectedUser).Return(errors.New("error from usecase")).Once()

		// Create the user controller and call the CreateUser method
		suite.controller.CreateUser(c)

		// Verify the response
		if w.Code != http.StatusInternalServerError {
			suite.T().Errorf("expected status code %d but got %d", http.StatusCreated, w.Code)
		}

		suite.usecase.AssertCalled(suite.T(), "CreateUser", mock.AnythingOfType("*domain.User"))

	})
}

func (suite *UserControllerSuite) TestDeleteUser() {
	suite.Run("success", func() {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID

		suite.usecase.On("DeleteUser", mock.Anything).Return(nil).Once()

		// Create the user controller and call the DeleteUser method
		suite.controller.DeleteUser(c)

		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}

		suite.usecase.AssertCalled(suite.T(), "DeleteUser", mock.Anything)
	})
	suite.Run("error", func() {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID

		suite.usecase.On("DeleteUser", mock.Anything).Return(errors.New("error from usecase")).Once()

		// Create the user controller and call the DeleteUser method
		suite.controller.DeleteUser(c)

		// Verify the response
		if w.Code != http.StatusInternalServerError {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}

		suite.usecase.AssertCalled(suite.T(), "DeleteUser", mock.Anything)
	})
	suite.Run("error_unauthorized", func() {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "user"}) // Example ObjectID

		suite.usecase.On("DeleteUser", mock.Anything).Return(nil).Once()

		// Create the user controller and call the DeleteUser method
		suite.controller.DeleteUser(c)

		// Verify the response
		if w.Code != http.StatusForbidden {
			suite.T().Errorf("expected status code %d but got %d", http.StatusUnauthorized, w.Code)
		}

		suite.usecase.AssertCalled(suite.T(), "DeleteUser", mock.Anything)
	})
}

func (suite *UserControllerSuite) TestGetUserByID() {

	suite.Run("sucess", func() {
		req, err := http.NewRequest(http.MethodGet, "/users/1", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID

		expectedUser := &domain.User{
			ID:    primitive.NewObjectID(),
			Name:  "test",
			Email: "teser@asd.cdds",
		}
		suite.usecase.On("GetUserByID", mock.Anything, mock.Anything).Return(expectedUser, nil).Once()
		// Create the user controller and call the GetUserByID method
		suite.controller.GetUserByID(c)
		// Verify the response
		if w.Code != http.StatusOK {

			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		var responseUser domain.User
		err = json.NewDecoder(w.Body).Decode(&responseUser)
		if err != nil {

			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(responseUser.ID, (*expectedUser).ID) {
			suite.T().Errorf("expected response body %+v but got %+v", *expectedUser, responseUser)
		}
		suite.usecase.AssertCalled(suite.T(), "GetUserByID", mock.Anything, mock.Anything)
	})
}

func (suite *UserControllerSuite) TestGetUsers() {
	suite.Run("sucess", func() {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			suite.T().Fatal(err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID

		ids := primitive.NewObjectID()
		users := []*domain.User{
			{
				ID:    ids,
				Name:  "user1",
				Email: "user1@example.com",
			},
			{
				ID:    ids,
				Name:  "user2",
				Email: "user2@example.com",
			},
		}
		suite.usecase.On("GetUsers").Return(users, nil).Once()
		// Create the user controller and call the GetUsers method
		suite.controller.GetUsers(c)
		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK,
				w.Code)
		}
		var responseUsers []*domain.User
		err = json.NewDecoder(w.Body).Decode(&responseUsers)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(responseUsers, users) {
			suite.T().Errorf("expected response body %+v but got %+v", users, responseUsers)
		}
		suite.usecase.AssertCalled(suite.T(), "GetUsers")
	})
}

func (suite *UserControllerSuite) TestUpdateUser() {
	// Mock the usecase method
	suite.Run("update success", func() {
		reqBody := `{"name":"testuser","email":""}`
		req, err := http.NewRequest(http.MethodPut, "/users/"+primitive.NewObjectID().Hex(), strings.NewReader(reqBody))
		if err != nil {
			suite.T().Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		// c.Set("id", primitive.NewObjectID().Hex())
		c.Params = gin.Params{
			{Key: "id", Value: primitive.NewObjectID().Hex()},
		}
		id := primitive.NewObjectID().Hex()
		c.Set("claims", &domain.Claims{UserID: id, Role: "admin"}) // Example ObjectID
		suite.usecase.On("UpdateUser", mock.Anything).Return(nil).Once()
		// Create the user controller and call the UpdateUser method
		suite.controller.UpdateUser(c)
		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		suite.usecase.AssertCalled(suite.T(), "UpdateUser", mock.Anything)
	})

}

func (suite *UserControllerSuite) TestLoginUser() {
	suite.Run("successfull login", func() {
		reqBody := `{"email":"abcd@ks.com","password":"testpassword"}`
		req, err := http.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
		if err != nil {
			suite.FailNow("error when creating the request")
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		// Prepare the expected user object and set up mock behavior
		hashedpass, err := infrastructure.HashPassword("testpassword")
		if err != nil {
			suite.T().Fatal(err)
		}
		expectedUser := domain.User{
			ID:       primitive.NewObjectID(),
			Name:     "testuser",
			Email:    "abcd@ks.com",
			Password: hashedpass,
			Role:     "user",
		}
		suite.usecase.On("Login", mock.Anything).Return(expectedUser, nil).Once()
		suite.usecase.On("GeneratesToken", mock.Anything).Return("token", nil)
		// Create the user controller and call the CreateUser method
		suite.controller.LoginUser(c)
		// Verify the response
		if w.Code != http.StatusOK {
			suite.T().Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}
		var responseUser domain.User
		err = json.NewDecoder(w.Body).Decode(&responseUser)
		if err != nil {
			suite.T().Fatal(err)
		}
		if !reflect.DeepEqual(responseUser.ID, (expectedUser).ID) {
			suite.T().Errorf("expected response body %+v but got %+v", expectedUser, responseUser)
		}
		suite.usecase.AssertCalled(suite.T(), "Login", mock.AnythingOfType("domain.User"))
	})
}

func ExecuteRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
