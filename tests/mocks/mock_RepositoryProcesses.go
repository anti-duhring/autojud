// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	processes "github.com/anti-duhring/autojud/internal/processes"
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryProcesses is an autogenerated mock type for the Repository type
type MockRepositoryProcesses struct {
	mock.Mock
}

type MockRepositoryProcesses_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoryProcesses) EXPECT() *MockRepositoryProcesses_Expecter {
	return &MockRepositoryProcesses_Expecter{mock: &_m.Mock}
}

// CreateProcessFollow provides a mock function with given fields: ctx, processID, userID
func (_m *MockRepositoryProcesses) CreateProcessFollow(ctx context.Context, processID string, userID string) (*processes.ProcessFollow, error) {
	ret := _m.Called(ctx, processID, userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateProcessFollow")
	}

	var r0 *processes.ProcessFollow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*processes.ProcessFollow, error)); ok {
		return rf(ctx, processID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *processes.ProcessFollow); ok {
		r0 = rf(ctx, processID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*processes.ProcessFollow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, processID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryProcesses_CreateProcessFollow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProcessFollow'
type MockRepositoryProcesses_CreateProcessFollow_Call struct {
	*mock.Call
}

// CreateProcessFollow is a helper method to define mock.On call
//   - ctx context.Context
//   - processID string
//   - userID string
func (_e *MockRepositoryProcesses_Expecter) CreateProcessFollow(ctx interface{}, processID interface{}, userID interface{}) *MockRepositoryProcesses_CreateProcessFollow_Call {
	return &MockRepositoryProcesses_CreateProcessFollow_Call{Call: _e.mock.On("CreateProcessFollow", ctx, processID, userID)}
}

func (_c *MockRepositoryProcesses_CreateProcessFollow_Call) Run(run func(ctx context.Context, processID string, userID string)) *MockRepositoryProcesses_CreateProcessFollow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRepositoryProcesses_CreateProcessFollow_Call) Return(_a0 *processes.ProcessFollow, _a1 error) *MockRepositoryProcesses_CreateProcessFollow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryProcesses_CreateProcessFollow_Call) RunAndReturn(run func(context.Context, string, string) (*processes.ProcessFollow, error)) *MockRepositoryProcesses_CreateProcessFollow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepositoryProcesses creates a new instance of MockRepositoryProcesses. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryProcesses(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryProcesses {
	mock := &MockRepositoryProcesses{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
