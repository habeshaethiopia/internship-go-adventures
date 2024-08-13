package usecases

import (
	"errors"
	domain "task/Domain"
	infrastructure "task/Infrastructure"
	"time"

	// "github.com/fatih/color"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	JwtService     infrastructure.JWTService
}

// GenerateToken implements domain.UserUsecase.

func NewUserUsecase(userRepository domain.UserRepository, jwtService infrastructure.JWTService) domain.UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
		JwtService:     jwtService,
	}
}

// Login implements domain.UserUsecase.
func (u *userUsecase) Login(user domain.User) (domain.User, error) {
	storedUser, err := u.UserRepository.GetUserByEmail(user.Email)
	return storedUser, err
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
	existingUser, err := u.UserRepository.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("user with email " + existingUser.Email + " already exists")
	}
	// Generate a unique ID for the user
	user.ID = primitive.NewObjectID()

	// Set the created and updated timestamps
	now := time.Now()
	user.Created_at = now
	user.Updated_at = now
	haskedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = haskedPassword
	// Save the user to the repository
	err = u.UserRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUser implements domain.UserUsecase.
func (u *userUsecase) DeleteUser(id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return u.UserRepository.DeleteUser(idObj)
}

// GetUserByID implements domain.UserUsecase.
func (u *userUsecase) GetUserByID(id string) (*domain.User, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.User{}, err
	}

	return u.UserRepository.GetUserByID(idObj)
}

// GetUsers implements domain.UserUsecase.
func (u *userUsecase) GetUsers() ([]*domain.User, error) {
	return u.UserRepository.GetUsers()
}

// UpdateUser implements domain.UserUsecase.
func (u *userUsecase) UpdateUser(user *domain.User) error {

	user.Updated_at = time.Now()

	// Save the user to the repository
	return u.UserRepository.UpdateUser(user)
}
func (u *userUsecase) GeneratesToken(claim domain.Claims) (string, error) {
	return u.JwtService.GenerateToken(&claim)
}
func (u *userUsecase) Get_secret_key() []byte {
	return u.JwtService.SecretKey
}
