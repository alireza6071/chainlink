// Code generated by mockery v2.21.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
)

// HeadTrackable is an autogenerated mock type for the HeadTrackable type
type HeadTrackable struct {
	mock.Mock
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *HeadTrackable) OnNewLongestChain(ctx context.Context, head *types.Head) {
	_m.Called(ctx, head)
}

type mockConstructorTestingTNewHeadTrackable interface {
	mock.TestingT
	Cleanup(func())
}

// NewHeadTrackable creates a new instance of HeadTrackable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHeadTrackable(t mockConstructorTestingTNewHeadTrackable) *HeadTrackable {
	mock := &HeadTrackable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
