package usecases

import (
	"errors"
	domain "task/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

// Login implements domain.UserUsecase.
func (u *userUsecase) Login(*domain.User) ([]byte, error) {
	panic("unimplemented")
}

// CreateUser implements domain.UserUsecase.
func (u *userUsecase) CreateUser(user *domain.User) error {
	// Validate user data
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}

	// Generate a unique ID for the user
	user.ID = primitive.NewObjectID()

	// Set the created and updated timestamps
	now := time.Now()
	user.Created_at = now
	user.Updated_at = now

	// Save the user to the repository
	err := u.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser implements domain.UserUsecase.
func (u *userUsecase) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUserByID implements domain.UserUsecase.
func (u *userUsecase) GetUserByID(id string) (*domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserUsecase.
func (u *userUsecase) GetUsers() ([]*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUsecase.
func (u *userUsecase) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
		contextTimeout: timeout,
	}
}
