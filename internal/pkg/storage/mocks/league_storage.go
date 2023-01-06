// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// LeagueStorage is an autogenerated mock type for the LeagueStorage type
type LeagueStorage struct {
	mock.Mock
}

type LeagueStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *LeagueStorage) EXPECT() *LeagueStorage_Expecter {
	return &LeagueStorage_Expecter{mock: &_m.Mock}
}

// GetLeagueByID provides a mock function with given fields: ctx, id
func (_m *LeagueStorage) GetLeagueByID(ctx context.Context, id int32) (model.League, error) {
	ret := _m.Called(ctx, id)

	var r0 model.League
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.League); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.League)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeagueStorage_GetLeagueByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLeagueByID'
type LeagueStorage_GetLeagueByID_Call struct {
	*mock.Call
}

// GetLeagueByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *LeagueStorage_Expecter) GetLeagueByID(ctx interface{}, id interface{}) *LeagueStorage_GetLeagueByID_Call {
	return &LeagueStorage_GetLeagueByID_Call{Call: _e.mock.On("GetLeagueByID", ctx, id)}
}

func (_c *LeagueStorage_GetLeagueByID_Call) Run(run func(ctx context.Context, id int32)) *LeagueStorage_GetLeagueByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *LeagueStorage_GetLeagueByID_Call) Return(_a0 model.League, _a1 error) *LeagueStorage_GetLeagueByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewLeagueStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewLeagueStorage creates a new instance of LeagueStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLeagueStorage(t mockConstructorTestingTNewLeagueStorage) *LeagueStorage {
	mock := &LeagueStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}