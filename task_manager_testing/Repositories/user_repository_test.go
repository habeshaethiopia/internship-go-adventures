package repositories

import (
	"errors"
	domain "task/Domain"
	mocks "task/mock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepositorySuite struct {
	suite.Suite
	userRepository domain.UserRepository
	mockCollection *mocks.Collection
	// mockCursor     *mocks.Cursor
	mockDatabase     *mocks.Database
	mockSingleResult *mocks.SingleResult
	mockcursor       *mocks.Cursor
}

func (suite *userRepositorySuite) SetupSuite() {
	suite.mockDatabase = new(mocks.Database)
	suite.mockCollection = new(mocks.Collection)
	suite.mockSingleResult = new(mocks.SingleResult)
	suite.mockcursor = new(mocks.Cursor)
	suite.mockDatabase.On("Collection", mock.Anything).Return(suite.mockCollection)
	suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult)
}

func (suite *userRepositorySuite) SetupTest() {
	suite.userRepository = NewUserRepository(suite.mockDatabase, "user")

}
func (suite *userRepositorySuite) TestCreateUser() {
	suite.Run("Success", func() {
		suite.mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil).Once()

		// Use the custom mock type for FindOne to return an error
		//
		suite.mockSingleResult.On("Decode", mock.Anything).Return(errors.New("not found")).Once()
		// suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult)
		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()
		user := domain.User{
			Email:    "124354",
			Password: "234",
		}
		err := suite.userRepository.CreateUser(&user)
		suite.NoError(err)
		suite.mockCollection.AssertCalled(suite.T(), "InsertOne", mock.Anything, mock.Anything)
	})
	suite.Run("Error", func() {
		suite.mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil).Once()
		suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Once()
		// suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult)
		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()
		user := domain.User{
			Email:    "",
			Password: "",
		}

		err := suite.userRepository.CreateUser(&user)
		suite.Error(err)
		suite.mockCollection.AssertCalled(suite.T(), "InsertOne", mock.Anything, mock.Anything)

	})

}
func (suite *userRepositorySuite) TestGetUserByEmail() {
	suite.Run("Success", func() {
		user := domain.User{
			Email:    "124354",
			Password: "234",
		}
		suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			*arg = user
		}).Once()

		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()

		result, err := suite.userRepository.GetUserByEmail(user.Email)
		suite.Equal(result, user)
		suite.NoError(err)
		suite.mockCollection.AssertCalled(suite.T(), "FindOne", mock.Anything, mock.Anything)
	})

}
func (suite *userRepositorySuite) TestDeleteUser() {
	suite.Run("Success", func() {
		suite.mockCollection.On("DeleteOne", mock.Anything, mock.Anything).Return(int64(1), nil).Once()
		id := primitive.NewObjectID()
		err := suite.userRepository.DeleteUser(id)
		suite.NoError(err)
		suite.mockCollection.AssertCalled(suite.T(), "DeleteOne", mock.Anything, mock.Anything)
	})
	suite.Run("Error", func() {
		suite.mockCollection.On("DeleteOne", mock.Anything, mock.Anything).Return(int64(0), errors.New("error")).Once()
		id := primitive.NewObjectID()
		err := suite.userRepository.DeleteUser(id)
		suite.Error(err)
		suite.mockCollection.AssertCalled(suite.T(), "DeleteOne", mock.Anything, mock.Anything)
	})
}
func (suite *userRepositorySuite) TestGetUserByID() {
	suite.Run("Success", func() {
		user := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "124354",
			Password: "234",
		}
		suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			*arg = user
		}).Once()

		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()

		result, err := suite.userRepository.GetUserByID(user.ID)
		suite.Equal(*result, user)
		suite.NoError(err)
		suite.mockCollection.AssertCalled(suite.T(), "FindOne", mock.Anything, mock.Anything)
	})
	suite.Run("Error", func() {

		suite.mockSingleResult.On("Decode", mock.Anything).Return(errors.New("not found")).Once()
		// suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult)
		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()
		
		suite.mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()

		_, err := suite.userRepository.GetUserByID(primitive.NewObjectID())
		suite.Error(err)
		suite.mockCollection.AssertCalled(suite.T(), "FindOne", mock.Anything, mock.Anything)
	})


}
func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}
