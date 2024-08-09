package repositories

import (
	"task/Domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(user *domain.User) error {
	panic("unimplemented")
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUserByID implements domain.UserRepository.
func (u *userRepository) GetUserByID(id string) (*domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserRepository.
func (u *userRepository) GetUsers() ([]*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
