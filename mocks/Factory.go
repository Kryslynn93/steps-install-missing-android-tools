// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	command "github.com/bitrise-io/go-utils/v2/command"
	mock "github.com/stretchr/testify/mock"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, args, opts
func (_m *Factory) Create(name string, args []string, opts *command.Opts) command.Command {
	ret := _m.Called(name, args, opts)

	var r0 command.Command
	if rf, ok := ret.Get(0).(func(string, []string, *command.Opts) command.Command); ok {
		r0 = rf(name, args, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(command.Command)
		}
	}

	return r0
}

type mockConstructorTestingTNewFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFactory(t mockConstructorTestingTNewFactory) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
