// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	bridges "github.com/smartcontractkit/chainlink/core/bridges"
)

// ExternalInitiatorManager is an autogenerated mock type for the ExternalInitiatorManager type
type ExternalInitiatorManager struct {
	mock.Mock
}

// DeleteJob provides a mock function with given fields: webhookSpecID
func (_m *ExternalInitiatorManager) DeleteJob(webhookSpecID int32) error {
	ret := _m.Called(webhookSpecID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(webhookSpecID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindExternalInitiatorByName provides a mock function with given fields: name
func (_m *ExternalInitiatorManager) FindExternalInitiatorByName(name string) (bridges.ExternalInitiator, error) {
	ret := _m.Called(name)

	var r0 bridges.ExternalInitiator
	if rf, ok := ret.Get(0).(func(string) bridges.ExternalInitiator); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bridges.ExternalInitiator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notify provides a mock function with given fields: webhookSpecID
func (_m *ExternalInitiatorManager) Notify(webhookSpecID int32) error {
	ret := _m.Called(webhookSpecID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(webhookSpecID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewExternalInitiatorManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewExternalInitiatorManager creates a new instance of ExternalInitiatorManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExternalInitiatorManager(t mockConstructorTestingTNewExternalInitiatorManager) *ExternalInitiatorManager {
	mock := &ExternalInitiatorManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
