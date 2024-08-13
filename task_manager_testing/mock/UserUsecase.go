// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "task/Domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserUsecase) CreateUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserUsecase) DeleteUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GeneratesToken provides a mock function with given fields: claim
func (_m *UserUsecase) GeneratesToken(claim domain.Claims) (string, error) {
	ret := _m.Called(claim)

	if len(ret) == 0 {
		panic("no return value specified for GeneratesToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Claims) (string, error)); ok {
		return rf(claim)
	}
	if rf, ok := ret.Get(0).(func(domain.Claims) string); ok {
		r0 = rf(claim)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.Claims) error); ok {
		r1 = rf(claim)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *UserUsecase) GetUsers() ([]*domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get_secret_key provides a mock function with given fields:
func (_m *UserUsecase) Get_secret_key() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get_secret_key")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Login provides a mock function with given fields: u
func (_m *UserUsecase) Login(u domain.User) (domain.User, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (domain.User, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *UserUsecase) UpdateUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
