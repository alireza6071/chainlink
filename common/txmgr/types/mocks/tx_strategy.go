// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TxStrategy is an autogenerated mock type for the TxStrategy type
type TxStrategy struct {
	mock.Mock
}

// PruneQueue provides a mock function with given fields: ctx, pruneService
func (_m *TxStrategy) PruneQueue(ctx context.Context, pruneService types.UnstartedTxQueuePruner) (int64, error) {
	ret := _m.Called(ctx, pruneService)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.UnstartedTxQueuePruner) (int64, error)); ok {
		return rf(ctx, pruneService)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.UnstartedTxQueuePruner) int64); ok {
		r0 = rf(ctx, pruneService)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.UnstartedTxQueuePruner) error); ok {
		r1 = rf(ctx, pruneService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subject provides a mock function with given fields:
func (_m *TxStrategy) Subject() uuid.NullUUID {
	ret := _m.Called()

	var r0 uuid.NullUUID
	if rf, ok := ret.Get(0).(func() uuid.NullUUID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uuid.NullUUID)
	}

	return r0
}

type mockConstructorTestingTNewTxStrategy interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxStrategy creates a new instance of TxStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxStrategy(t mockConstructorTestingTNewTxStrategy) *TxStrategy {
	mock := &TxStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
