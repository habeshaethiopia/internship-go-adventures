// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "task/Domain"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *TaskRepository) CreateTask(task *domain.Task) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTask provides a mock function with given fields: id
func (_m *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskByID provides a mock function with given fields: id
func (_m *TaskRepository) GetTaskByID(id primitive.ObjectID) (*domain.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (*domain.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields:
func (_m *TaskRepository) GetTasks() ([]*domain.Task, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTasks")
	}

	var r0 []*domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: task
func (_m *TaskRepository) UpdateTask(task *domain.Task) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
