package usecases

import (
	"errors"
	"fmt"
	domain "task/Domain"
	infrastructure "task/Infrastructure"
	mocks "task/mock"
	"testing"

	// Add this import to use the AnythingOfType function
	"github.com/fatih/color"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUseCaseSuite struct {
	suite.Suite
	userUseCase  domain.UserUsecase   // Use case under test
	mockUserRepo mocks.UserRepository // Mocked repository
}

func (suite *userUseCaseSuite) SetupSuite() {

	color.Green("Start user Usecase Test")
}

// SetupTest sets up the test environment
func (suite *userUseCaseSuite) SetupTest() {
	// Create a new mock repository
	suite.userUseCase = NewUserUsecase(&suite.mockUserRepo, *infrastructure.NewJWTService("test"))

}
func (suite *userUseCaseSuite) TearDownSuite() {
	color.Green("user usecase is done")
	// No need to clean up in this case
}
func (suite *userUseCaseSuite) TestCreateUser() {
	suite.mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil).Once()
	suite.mockUserRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(domain.User{}, errors.New("no user found")).Once()

	
	user := domain.User{
		Name:     "test",
		Password: "test",
		Email:    "test@example.com",
		Role:     "user",
	}
	err := suite.userUseCase.CreateUser(&user)

	suite.NoError(err)
	suite.mockUserRepo.AssertCalled(suite.T(), "CreateUser", mock.AnythingOfType("*domain.User"))

}
func (suite *userUseCaseSuite) TestCreateUserAlreadyFound() {
	suite.mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil).Once()
	suite.mockUserRepo.On("GetUserByEmail", "test@example.com").Return(domain.User{},nil).Once()

	
	user := domain.User{
		Name:     "test",
		Password: "test",
		Email:    "test@example.com",
		Role:     "user",
	}
	err := suite.userUseCase.CreateUser(&user)

	suite.Error(err)
	suite.mockUserRepo.AssertCalled(suite.T(), "GetUserByEmail", "test@example.com")
	suite.mockUserRepo.AssertNotCalled(suite.T(), "CreateUser",&user)


}
func (suite *userUseCaseSuite) TestGetUser() {
	suite.mockUserRepo.On("GetUsers").Return([]*domain.User{}, nil).Once()
	result, err := suite.userUseCase.GetUsers()
	suite.NoError(err)
	suite.NotNil(result)
	suite.mockUserRepo.AssertCalled(suite.T(), "GetUsers")
}

func (s *userUseCaseSuite) TestGetUserByID() {
	s.mockUserRepo.On("GetUserByID", mock.AnythingOfType("primitive.ObjectID")).Return(&domain.User{}, nil).Once()
	result, err := s.userUseCase.GetUserByID(primitive.NewObjectID().Hex())
	s.NoError(err)
	s.NotNil(result)
	s.mockUserRepo.AssertCalled(s.T(), "GetUserByID", mock.AnythingOfType("primitive.ObjectID"))
}

func (s *userUseCaseSuite) TestGetUserByIDnotfound() {
	s.mockUserRepo.On("GetUserByID", mock.AnythingOfType("primitive.ObjectID")).Return(nil, fmt.Errorf("not found")).Once()
	result, err := s.userUseCase.GetUserByID(primitive.NewObjectID().Hex())
	s.Error(err)
	s.Nil(result)
	s.mockUserRepo.AssertCalled(s.T(), "GetUserByID", mock.AnythingOfType("primitive.ObjectID"))
}
func (s *userUseCaseSuite) TestDeleteUser() {
	s.mockUserRepo.On("DeleteUser", mock.AnythingOfType("primitive.ObjectID")).Return(nil).Once()
	err := s.userUseCase.DeleteUser(primitive.NewObjectID().Hex())
	s.NoError(err)
	s.mockUserRepo.AssertCalled(s.T(), "DeleteUser", mock.AnythingOfType("primitive.ObjectID"))
}
func (s *userUseCaseSuite) TestDeleteUserError() {
	s.mockUserRepo.On("DeleteUser", mock.AnythingOfType("primitive.ObjectID")).Return(fmt.Errorf("error")).Once()
	err := s.userUseCase.DeleteUser(primitive.NewObjectID().Hex())
	s.Error(err)
	s.mockUserRepo.AssertCalled(s.T(), "DeleteUser", mock.AnythingOfType("primitive.ObjectID"))
}
func (s *userUseCaseSuite) TestUpdateUser() {
	s.mockUserRepo.On("UpdateUser", mock.AnythingOfType("*domain.User")).Return(nil).Once()
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "test",
		Password: "somepasss",
	}
	err := s.userUseCase.UpdateUser(&user)
	s.NoError(err)
	s.mockUserRepo.AssertCalled(s.T(), "UpdateUser", mock.AnythingOfType("*domain.User"))
}
func (s *userUseCaseSuite) TestUpdateUserNotFound() {
	s.mockUserRepo.On("UpdateUser", mock.AnythingOfType("*domain.User")).Return(fmt.Errorf("not found")).Once()
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "test",
		Password: "somepasss",
	}
	err := s.userUseCase.UpdateUser(&user)
	s.Error(err)
	s.mockUserRepo.AssertCalled(s.T(), "UpdateUser", mock.AnythingOfType("*domain.User"))
}
func (s *userUseCaseSuite) TestUpdateUserError() {
	s.mockUserRepo.On("UpdateUser", mock.AnythingOfType("*domain.User")).Return(fmt.Errorf("error")).Once()
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "test",
		Password: "somepass",
	}
	err := s.userUseCase.UpdateUser(&user)
	s.Error(err)
	s.mockUserRepo.AssertCalled(s.T(), "UpdateUser", mock.AnythingOfType("*domain.User"))
}
func (s *userUseCaseSuite) TestLogin() {
	s.mockUserRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(domain.User{}, nil).Once()
	result, err := s.userUseCase.Login(domain.User{
		Email:    "your_email",
		Password: "some string",
	})
	s.NoError(err)
	s.NotNil(result)
	s.mockUserRepo.AssertCalled(s.T(), "GetUserByEmail", mock.AnythingOfType("string"))
}
func (s *userUseCaseSuite) TestLoginError() {
	s.mockUserRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(domain.User{}, fmt.Errorf("error")).Once()
	result, err := s.userUseCase.Login(domain.User{
		Email:    "your_email",
		Password: "some string",
	})
	s.Error(err)
	s.Empty(result)
	s.mockUserRepo.AssertCalled(s.T(), "GetUserByEmail", mock.AnythingOfType("string"))
}
func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUseCaseSuite))
}
